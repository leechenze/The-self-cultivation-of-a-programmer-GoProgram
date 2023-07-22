package main

import (
	pb "GoMicro/GRPC/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	// 创建与gRPC服务端的连接
	dial, err := grpc.Dial("127.0.0.1:1024", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常: %v\n", err)
	}
	// 关闭连接
	defer dial.Close()
	// 实例化gRPC客户端
	client := pb.NewUserInfoServiceClient(dial)
	// 组装参数
	req := new(pb.UserRequest)
	req.Name = "trump"
	// 调用接口
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常: %v\n", err)
	}
	fmt.Printf("响应结果: %v\n", resp)
}
