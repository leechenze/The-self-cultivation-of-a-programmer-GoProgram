package goroutines

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg, %v\n", msg)
		// 没循环一次睡眠100毫秒
		time.Sleep(time.Millisecond * 100)
	}
}

func GoRoutines() {
	println("========================Go Routines========================")
	println()

	// 通过 go 关键字的声明指定启动一个协程来同步执行此程序。
	go showMsg("java")
	println()
	showMsg("golang")

	println()
	println("========================Go Routines========================")

}
