package goChanIterate

import "fmt"

var channel = make(chan int)

func GoChanIterate() {
	println("========================Channel Iterate========================")
	println()

	// 设置channel数据
	go func() {
		for i := 0; i < 2; i++ {
			channel <- i
		}
		// 关闭通道
		close(channel)
	}()

	// 读取channel数据
	for i := 0; i < 3; i++ {
		readChannel := <-channel
		// 如果不关闭通道，通道中只有两个值，在读取第三个时就会报错。
		fmt.Printf("readChannel: %v\n", readChannel)
	}

	/** range 遍历通道 */
	for val := range channel {
		fmt.Printf("readChannel: %v \n", val)
	}

	println()
	println("========================Channel Iterate========================")
}
