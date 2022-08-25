package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	pb "go-micro/01/proto"
)

// Hello 声明结构体
type Hello struct {
}

func (g *Hello) Info(ctx context.Context, req *pb.InfoRequest, rep *pb.InfoResponse) error {
	rep.Msg = "你好" + req.Username
	return nil
}

func main() {
	// 1.服务端实力
	service := micro.NewService(
		// 设置微服务的名，用来访问
		// micro命令: micro call hello Hello.Info {\"username\":\"zhangsan\"}
		micro.Name("hello"),
	)

	// 2.初始化
	service.Init()
	// 3.服务注册
	err := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	err = service.Run()
	if err != nil {
		fmt.Println(err)
	}
}
