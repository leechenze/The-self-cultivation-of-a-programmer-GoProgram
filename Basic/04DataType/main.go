package main

import "fmt"

func function() {}

func main() {

	println("========================Basic Type start========================")

	println()
	println("基本类型")
	name := "tom"
	age := 20
	isMarried := true
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)

	println()
	println("指针类型")
	a := 10
	b := &a
	fmt.Printf("%T\n", b)

	println()
	println("数组类型")
	arr := []int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr)
	fmt.Printf("%v\n", arr)

	println()
	println("函数类型")
	fmt.Printf("%T\n", function)
	fmt.Println(function)

	println()
	println()
	println()
	println("========================Basic Type end========================")

}
