package slice

import "fmt"

func Slice() {
	println("========================Slice module console========================")
	println()

	// 切片的声明和数组类似
	var a1 []int
	var a2 []string
	fmt.Printf("a1 %v\n", a1)
	fmt.Printf("a2 %v\n", a2)
	// 通过make创建切片
	var a3 = make([]int, 2)
	fmt.Printf("a3 %v\n", a3)
	fmt.Printf("a3[1] %v\n", a3[1])
	// 切片初始化
	var a4 = []int{1, 2, 3, 4}
	fmt.Printf("len(a4) %v\n", len(a4))
	fmt.Printf("cap(a4) %v\n", cap(a4))

	println()
	println("========================Slice module console========================")
}
