package main

func main() {
	println("=========================Variable start=========================")
	var name string
	println("name: ", name)
	var age int
	println("age:", age)
	var isMarried bool
	println("isMarried:", isMarried)

	println()
	// 批量声明
	var (
		name1      string
		age1       int
		isMarried1 bool
	)
	println("name1:", name1)
	println("age1:", age1)
	println("isMarried1:", isMarried1)

	println()
	// 变量初始化
	var name2 string = "tom"
	println("name2: ", name2)
	var age2 int = 90
	println("age2:", age2)
	var isMarried2 bool = true
	println("isMarried2:", isMarried2)

	println()
	// 类型推导
	// 在声明变量时，可以根据初始值进行类型推导，从而省略类型。
	var name3 = "Ken Thompson"
	println("name3: ", name3)
	var age3 = 88
	println("age3:", age3)
	var isMarried3 = true
	println("isMarried3:", isMarried3)

	println()
	// 批量初始化
	var name4, age4, isMarried4 = "lincoln", 222, true
	println("name4: ", name4)
	println("age4:", age4)
	println("isMarried4:", isMarried4)

	println()
	// 短变量声明
	// 在函数内部，可以使用 := 运算符对变量进行声明和初始化，这样可以即省略变量又省略了类型（:=会自动进行类型推导）
	name5 := "trump"
	println(name5)

	println()
	// 匿名变量
	name6, age6 := getNameAndAge()
	println("name6: ", name6)
	println("age6:", age6)

	println("=========================Variable end=========================")
	println("=========================Constant start=========================")

	println()
	const PI float32 = 3.1
	// 类型推导
	const PI1 = 3.1
	println(PI)
	println(PI1)

	// 批量声明
	println()
	const (
		A = 100
		B = 200
	)
	println(A, B)

	println()
	const A1, B1 = 300, 400
	println(A1, B1)

	println()
	// 特殊常量值：iota
	// iota比较特殊，可以被认为是一个可被编译器修改的常量，它默认的初始值是0，每调用一次加1，遇到const关键字时被重置为0。
	const (
		I1 = iota
		I2 = iota
		I3 = iota
	)
	println(I1, I2, I3)

	println()
	// 使用 _ 跳过某些值
	const (
		I11 = iota
		_
		I22 = iota
	)
	println(I11, I22)

	println()
	// 使用 _ 跳过某些值
	const (
		I111 = iota
		I222 = 999
		I333 = iota
	)
	println(I111, I222, I333)

	println()
	println("=========================Constant end=========================")

}

// 匿名变量：(name string, age int) 省略为 (string, int)
func getNameAndAge() (string, int) {
	return "clinton", 20
}
