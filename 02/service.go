package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/errors"
	pb "go-micro/02/proto"
	"log"
)

type Example struct {
}

type Foo struct {
}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Print("收到Example.Call请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no name")
	}
	rsp.Message = "Call.CallFunc接收到了你的请求" + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *pb.EmptyRequest, rsp *pb.EmptyResponse) error {
	log.Print("Foo.Bar")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	service.Init()
	err := pb.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		fmt.Println(err)
	}

	err = pb.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		fmt.Println(err)
	}
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
