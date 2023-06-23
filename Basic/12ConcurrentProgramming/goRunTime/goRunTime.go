package goRunTime

import (
	"runtime"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		println(msg)
		if i == 2 {
			// 当执行到第二轮循环时退出当前协程（后续的msg将不会再打印）
			runtime.Goexit()
		}
	}
}

/** GOMAXPROCS 和 NumCPU */
func a() {
	for i := 0; i < 10; i++ {
		println("ai", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		println("bi", i)
	}
}

func GoRunTime() {
	println("========================Go Routines========================")
	println()

	go showMsg("java")
	// 让出当前协程的CPU时间片，让其他协程先执行
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		println("golang")
	}
	println("end")

	/** GOMAXPROCS 和 NumCPU */
	println()
	println("GOMAXPROCS 和 NumCPU")
	// 查看CPU的核心数
	println(runtime.NumCPU())
	// 设置CPU的最大核心数，声明分别有两个CPU核心对a和b函数进行交替执行。如果设置一个核心的话，那么就是a执行完后才能执行b。
	runtime.GOMAXPROCS(2)
	go a()
	// 执行Gosched，以等待a和b的运行结束
	runtime.Gosched()
	go b()
	// 执行Gosched，以等待a和b的运行结束
	runtime.Gosched()

	println()
	println("========================Go Routines========================")

}
