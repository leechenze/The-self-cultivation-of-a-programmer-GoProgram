package goSelect

import (
	"fmt"
	"time"
)

var chanInt = make(chan int, 0)
var chanStr = make(chan string)

func GoSelect() {
	println("========================Concurrent Programming========================")
	println()

	go func() {
		chanInt <- 100
		chanStr <- "hello"
		defer close(chanInt)
		defer close(chanStr)
	}()

	for {
		select {
		case readChanInt := <-chanInt:
			fmt.Printf("readChanInt: %v\n", readChanInt)
		case readChanStr := <-chanStr:
			fmt.Printf("readChanStr: %v\n", readChanStr)
		default:
			println("default")
		}
		time.Sleep(time.Second)
	}

	println()
	println("========================Concurrent Programming========================")

}
