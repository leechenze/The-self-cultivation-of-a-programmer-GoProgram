package main

import (
	"pointerAndStruct/arrayPointer"
	"pointerAndStruct/basicPointer"
	"pointerAndStruct/structInit"
	"pointerAndStruct/structMethod"
	"pointerAndStruct/structPointer"
)

func main() {
	println("========================Pointer And Struct========================")
	// 基础指针
	basicPointer.BasicPointer()
	// 数组指针
	arrayPointer.ArrayPointer()
	// 结构体指针
	structPointer.StructPointer()
	// 结构体
	structInit.StructInit()
	// 结构体方法
	structMethod.StructMethod()

}
