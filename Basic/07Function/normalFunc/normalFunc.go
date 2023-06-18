package normalFunc

import "fmt"

// 普通函数定义
func sum(a int, b int) (res int) {
	res = a + b
	return a + b
}

// 或
func sum1(a int, b int) int {
	return a + b
}

// 多个返回值
func mutilRet() (string, int) {
	return "leechenze", 20
}

// 多个返回值定义返回名称时，可以直接return
func mutilRet1() (name string, age int) {
	name = "leechenze"
	age = 20
	return // 等价于 return "leechenze", 20
}

// 函数变长参数
func mutilParams(args ...int) {
	for key, val := range args {
		fmt.Printf("key %v\n", key)
		fmt.Printf("val %v\n", val)
	}
}

func NormalFunc() {
	println("========================Normal Func========================")
	println()

	res := sum(1, 2)
	fmt.Printf("res %v\n", res)

	// 多返回值
	println()
	// res1, res2 := mutilRet()
	res1, res2 := mutilRet1()
	fmt.Printf("res1 %v\n", res1)
	fmt.Printf("res2 %v\n", res2)

	// 函数变长参数
	mutilParams(11, 22, 33, 44, 55)

	// 函数类型声明
	println()
	println("函数类型声明")
	type fn func(int, int) int
	var foo fn
	foo = sum
	fmt.Printf("foo(2,3) %v\n", foo(2, 3))

	println()
	println("========================Normal Func========================")
}
