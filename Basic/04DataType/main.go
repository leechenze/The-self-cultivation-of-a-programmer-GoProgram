package main

import (
	"dataType/basicType"
	"dataType/boolean"
	"dataType/formatOutput"
	"dataType/numberType"
	"dataType/stringType"
)

func main() {
	// 基础数据类型
	basicType.BasicTypeHelloWorld()
	// 布尔类型
	boolean.BoolTypeHelloWorld()
	// 数字类型
	numberType.NumberType()
	// 字符串
	stringType.StringType()
	// 格式化输出
	formatOutput.FormatOutput()
}
