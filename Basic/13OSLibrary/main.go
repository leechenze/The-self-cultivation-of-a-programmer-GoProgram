package main

import (
	"osLibrary/goFileDir"
	"osLibrary/goProce"
)

func main() {
	println("========================OS Library========================")
	// 文件目录操作
	goFileDir.GoFileDir()
	// 进程相关操作
	goProce.GoProce()

}
