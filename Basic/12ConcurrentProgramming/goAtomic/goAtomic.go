package goAtomic

import (
	"fmt"
	"sync/atomic"
)

func demoAddSub() {
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, -1)
	fmt.Printf("i: %v\n", i)

	var j int64 = 100
	atomic.AddInt64(&j, 1)
	fmt.Printf("j: %v\n", j)
	atomic.AddInt64(&j, -1)
	fmt.Printf("j: %v\n", j)
}

func demoLoadStore() {
	var i int32 = 100
	// 载入（读操作）
	fmt.Printf("atomic.LoadInt32(&i): %v\n", atomic.LoadInt32(&i))
	fmt.Printf("i: %v\n", i)
	// 存储（写操作）
	atomic.StoreInt32(&i, 200)
	fmt.Printf("i: %v\n", i)
}

func demoCAS() {
	var i int32 = 100
	// 将 i 重新从200新值 替换掉100的旧值
	ok := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("ok: %v\n", ok)
	fmt.Printf("i: %v\n", i)
}

func GoAtomic() {
	println("========================Concurrent Programming========================")
	println()

	/** 原子操作加减演示 */
	println()
	demoAddSub()

	/** 原子操作载入和存储演示 */
	println()
	demoLoadStore()

	/** CAS操作演示 */
	println()
	demoCAS()

	// CAS（compare and swap）是其他操作的根基，除此之外还有一个Swap的操作
	// 这个操作比较暴力，没有对比操作，直接替换，一般用到的场景不多。

	println()
	println("========================Concurrent Programming========================")
}
