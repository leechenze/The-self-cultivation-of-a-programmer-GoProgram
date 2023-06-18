package main

import (
	"function/advancedFunc"
	"function/anonymous"
	"function/normalFunc"
)

func main() {
	println("========================Function========================")
	// 普通函数
	normalFunc.NormalFunc()
	// 高阶函数
	advancedFunc.AdvancedFunc()
	// 匿名函数
	anonymous.Anonymous()

}
