package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 声明矩形结构体
type Rect struct {
}

// 声明参数结构体
type Params struct {
	Width, Height int
}

// 定义求矩形面积的方法
func (r *Rect) Area(p Params, res *int) error {
	*res = p.Width * p.Height
	return nil
}

// 定义求矩形周长的方法
func (r *Rect) Perimeter(p Params, res *int) error {
	*res = (p.Width + p.Height) * 2
	return nil
}

// 普通RPC调用编码方式
func main1() {
	// 注册服务
	rect := new(Rect)
	rpc.Register(rect)
	// 把服务处理绑定到HTTP协议上
	rpc.HandleHTTP()
	// 监听服务被调用 Area 和 Perimeter 方法
	err := http.ListenAndServe(":1024", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// JsonRPC调用编码方式，作用：JSON是所有编程语言都通用的，所以JsonRPC可以使Go程序的RPC调用跨语言。
func main() {
	// 注册服务
	rect := new(Rect)
	rpc.Register(rect)
	// 监听服务被调用 Area 和 Perimeter 方法
	listen, err := net.Listen("tcp", "127.0.0.1:1024")

	if err != nil {
		log.Fatalf("Connect Error: %s \n", err)
	}
	// 循环监听服务
	for {
		connect, err1 := listen.Accept()
		if err1 != nil {
			// 如果有错误直接跳过当次循环即可，因为还需要监听其他的。
			continue
		}
		// 协程
		go func(conn net.Conn) {
			fmt.Println("new Client：监听到与客户端建立的一个连接")
			jsonrpc.ServeConn(conn)
		}(connect)
	}

}
