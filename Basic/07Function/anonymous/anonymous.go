package anonymous

import "fmt"

func Anonymous() {
	println("========================Anonymous Func========================")
	println()

	// 匿名函数初始化
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	res := max(1, 2)
	fmt.Printf("res %v\n", res)

	// 自调用函数
	res1 := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(2, 3)
	fmt.Printf("res1 %v\n", res1)

	// 闭包

	println()
	println("========================Anonymous Func========================")
}
