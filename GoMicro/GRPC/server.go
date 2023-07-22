package main

import (
	pb "GoMicro/GRPC/proto"
	"context"
	"fmt"
	grpc "google.golang.org/grpc"
	"net"
)

// 定义服务端实现约定的接口
type UserInfoService struct {
}

var u = UserInfoService{}

// 实现服务端需要实现的接口
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	// 在数据库查用户信息(模拟)
	if name == "trump" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"Sing", "Run"},
		}
	}
	err = nil
	return
}

func main() {
	// 监听
	addr := "127.0.0.1:1024"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常：%v\n", err)
	}
	fmt.Printf("开始监听:%v\n", addr)
	// 实例化gRPC
	server := grpc.NewServer()
	// 在gRPC上注册微服务
	pb.RegisterUserInfoServiceServer(server, &u)
	// 启动gRPC服务端
	server.Serve(listen)
}
