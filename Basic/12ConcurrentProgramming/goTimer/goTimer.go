package goTimer

import (
	"fmt"
	"time"
)

func GoTimer() {
	println("========================Concurrent Programming========================")
	println()

	// NewTimer
	// timer := time.NewTimer(time.Second * 2)
	// fmt.Printf("time.Now(): %v\n", time.Now())
	// t1 := <-timer.C // C表示channel，会进行阻塞指定的时间。
	// println("阻塞间隔时间为两秒")
	// // t1 就是返回的当前两秒后的时间，不指定 t1变量 接收阻塞后的时间一样可以阻塞。
	// fmt.Printf("t1: %v\n", t1)

	// After
	println()
	fmt.Printf("time.Now(): %v\n", time.Now())
	// <-time.After(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())

	// Sleep
	println()
	fmt.Printf("time.Now(): %v\n", time.Now())
	// time.Sleep(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())

	// Stop
	// println()
	// timer1 := time.NewTimer(time.Second * 2)
	// go func() {
	// 	<-timer1.C
	// 	println("timer1")
	// }()
	// // 可以将以下代码注释，看效果
	// s := timer1.Stop()
	// if s {
	// 	println("stop ...")
	// }
	//
	// time.Sleep(time.Second * 3)
	// println("main end ...")

	// Reset
	println()
	println("reset before")
	timer2 := time.NewTimer(time.Second * 5)
	timer2.Reset(time.Second)
	<-timer2.C
	println("reset after")

	println()
	println("========================Concurrent Programming========================")
}
