package gochannel

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建int类型的通道，只能传入int类型
var gochannel = make(chan int)

func send() {
	// 指定一个 10 范围之内的随机值，value就是要发送的数据
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("value of send %v\n", value)
	time.Sleep(time.Second * 5)
	// 将value发送到通道中。
	gochannel <- value
}

func GoChannel() {
	println("========================Go Routines========================")
	println()

	/** 从通道接收值 */
	// 使用defer，最后执行关闭通道
	defer close(gochannel)
	go send()
	println("wait ...")
	// 接收通道发送的数据
	value := <-gochannel
	fmt.Printf("value of main, %v\n", value)
	println("cahnnel end")

	println()
	println("========================Go Routines========================")
}
