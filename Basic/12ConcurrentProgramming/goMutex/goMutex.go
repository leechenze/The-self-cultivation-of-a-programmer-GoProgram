package goMutex

import (
	"fmt"
	"sync"
	"time"
)

var n int = 100
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	// Lock 和 Unlock 被锁定的协程，只有执行完之后，其他的协程才能切进来。
	lock.Lock()
	n += 1
	fmt.Printf("i++: %v\n", n)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}
func sub() {
	defer wg.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	n -= 1
	fmt.Printf("i--: %v\n", n)
	lock.Unlock()
}

func GoMutex() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()
	fmt.Printf("end n: %v\n", n)

}
