package entity

import "go-micro/demo/src/share/pb"

// User 结构体定义
type User struct {
	Id      int32  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	Phone   string `json:"phone" db:"Phone"`
}

func (u User) ToProtoUser() *pb.User {
	return &pb.User{
		Id:      u.Id,
		Name:    u.Name,
		Address: u.Address,
		Phone:   u.Phone,
	}
}
