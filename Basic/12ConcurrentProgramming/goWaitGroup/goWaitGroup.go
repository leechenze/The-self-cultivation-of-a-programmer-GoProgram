package goWaitGroup

import (
	"fmt"
	"sync"
)

// 声明waitGroup
var wg sync.WaitGroup

func showMsg(i int) {
	// 加个defer 保证后面输出之后才执行 wg.Done()
	// Done函数中就是执行的 wg.Add(-1)，表示当前协程执行完毕（退出当前子协程的同步等待）
	defer wg.Done()
	fmt.Printf("i %v\n", i)
}

func GoWaitGroup() {
	println("========================Go Routines========================")
	println()

	// 启动十个协程执行 showMsg
	for i := 0; i < 10; i++ {
		// 声明当前协程开始执行（进入当前子协程的同步等待）
		wg.Add(1)
		go showMsg(i)
	}

	// 执行等待（阻塞子协程的执行结束）
	wg.Wait()

	// 当main主协程执行完毕后，会将showMsg这些子协程没有执行完成的情况下的都杀死，导致子协程的程序没有执行。
	// 似此情况，就需要主协程（main）等待子协程（showMsg）的运行结束，就需要用到waitgroup
	println("end")

	println()
	println("========================Go Routines========================")
}
