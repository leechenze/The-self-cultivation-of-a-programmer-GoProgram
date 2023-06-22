package main

import (
	"interface/basicInterface"
	"interface/interfaceNest"
)

func main() {
	println("========================Interface========================")
	// 接口基本用法
	basicInterface.BasicInterface()
	// 接口嵌套
	interfaceNest.InterfaceNest()

}
