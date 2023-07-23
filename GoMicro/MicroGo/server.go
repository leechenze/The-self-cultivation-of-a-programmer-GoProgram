package main

import (
	pb "GoMicro/MicroGo/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

// 声明结构体
type Hello struct{}

// 实现接口方法
func (g *Hello) Info(ctx context.Context, req *pb.InfoRequest, res *pb.InfoResponse) error {
	res.Msg = "你好 " + req.Username
	return nil
}

func main() {
	// 得到微服务实例
	service := micro.NewService(
		// 设置微服务的名字为hello，用来做访问用的
		micro.Name("hello"),
	)

	// 初始化
	service.Init()

	// 服务注册
	err := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Printf("注册异常：%v\n", err)
	}

	// 启动微服务
	err = service.Run()
	if err != nil {
		fmt.Printf("启动异常：%v\n", err)
	}

}
