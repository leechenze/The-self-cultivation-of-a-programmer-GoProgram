package advancedFunc

import (
	"fmt"
)

// 函数作为参数
func sayHello(name string) {
	fmt.Printf("Hello, %v\n", name)
}
func func1(name string, f1 func(string)) {
	f1(name)
}

// 函数作为返回值
func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func calculate(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func AdvancedFunc() {
	println("========================Advanced Func========================")
	println()

	println("函数作为参数")
	func1("leechenze", sayHello)

	println()
	println("函数作为返回值")
	addOrSub := calculate("+")
	res := addOrSub(1, 2)
	// OR
	simpleRes := calculate("+")(1, 2)
	fmt.Printf("res %v\n", res)
	fmt.Printf("simpleRes %v\n", simpleRes)

	println()
	println("========================Advanced Func========================")
}
