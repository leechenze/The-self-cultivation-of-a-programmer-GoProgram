package anonymous

import "fmt"

// 闭包
func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包二
func calculate(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub
}

// 递归(死循环,别执行)
func recursion() {
	print("recursion is runing")
	recursion()
}

// 递归实现5的阶乘，正确结果为120
func fiveFactorial(a int) int {
	if a == 1 {
		return 1
	} else {
		return a * fiveFactorial(a-1)
	}
}

// 递归实现斐波那契数列
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func deferFunc() {
	println("defer func start runing")
	defer println("Part 3 Execution")
	defer println("Part 2 Execution")
	defer println("Part 1 Execution")
	println("defer func end runing")
}

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

	// 自执行函数
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
	println("闭包")
	f := add()
	res2 := f(10)
	fmt.Printf("res2 %v\n", res2)
	res3 := f(20)
	fmt.Printf("res3 %v\n", res3)
	res4 := f(30)
	fmt.Printf("res4 %v\n", res4)
	println()
	println("闭包二")
	add, sub := calculate(100)
	addRes := add(100)
	fmt.Printf("addRes %v\n", addRes)
	subRes := sub(50)
	fmt.Printf("subRes %v\n", subRes)

	// 递归
	// recursion()
	// 递归实现5的阶乘，正确结果120
	res5 := fiveFactorial(5)
	fmt.Printf("res5 %v\n", res5)
	// 递归实现斐波那契数列，正确结果为55
	res6 := fibonacci(10)
	fmt.Printf("res6 %v\n", res6)

	// defer
	println()
	println("defer")
	deferFunc()

	println()
	println("========================Anonymous Func========================")
}
