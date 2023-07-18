package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 声明参数结构体
type Params struct {
	Width, Height int
}

// 调用服务
func main() {
	// 连接远程RPC服务
	// http, err := rpc.DialHTTP("tcp", "127.0.0.1:1001")
	// JsonRPC连接远程RPC服务（客户端只需要改这一处即可）
	http, err := jsonrpc.Dial("tcp", "127.0.0.1:1001")
	if err != nil {
		log.Fatal(err)
	} else {
		// 定义结算结果的变量
		res := 0
		// 调用远程方法（求面积的方法）
		err1 := http.Call("Rect.Area", Params{Width: 50, Height: 50}, &res)
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Printf("面积是: %v\n", res)
		// 调用远程方法（求周长的方法）
		err2 := http.Call("Rect.Perimeter", Params{Width: 100, Height: 100}, &res)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Printf("周长是: %v\n", res)
	}
}
