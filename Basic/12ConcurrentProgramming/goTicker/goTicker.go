package goTicker

import (
	"fmt"
	"time"
)

func GoTicker() {
	println("========================Concurrent Programming========================")
	println()

	ticker := time.NewTicker(time.Second)

	/** Ticker基础演示 */
	// counter := 1
	// for _ = range ticker.C {
	// 	counter++
	// 	println("ticker")
	// 	if counter > 5 {
	// 		ticker.Stop()
	// 		break
	// 	}
	// }

	chanInt := make(chan int, 0)

	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for val := range chanInt {
		fmt.Printf("接收到的val是：%v\n", val)
		sum += val
		if sum >= 10 {
			ticker.Stop()
			break
		}
	}

	println()
	println("========================Concurrent Programming========================")
}
