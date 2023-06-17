package array

import "fmt"

func Array() {
	println("========================Array module console========================")
	println()

	// 数组声明，如果只是声明，没有初始化赋值不能用有等号（=）
	var a1 [2]int
	var a2 [3]string
	fmt.Printf("a1 %v\n", a1)
	fmt.Printf("a2 %v\n", a2)

	// 数组初始化，如果赋值的话不要忘记等号（=）
	println()
	a3 := [3]int{1, 2, 3}
	a4 := [5]int{1, 2, 3}
	a5 := [5]string{"hello", "world"}
	fmt.Printf("a3 %v\n", a3)
	fmt.Printf("a4 %v\n", a4)
	fmt.Printf("a5 %v\n", a5)

	// 省略数组的长度
	println()
	a6 := [...]int{1, 2, 3}
	fmt.Printf("a6 %v\n", a6)

	// 指定索引值的方式初始化
	a7 := [...]int{0: 11, 1: 22, 2: 33}
	fmt.Printf("a7 %v\n", a7)

	// 修改数组，数组越界问题，如果数组声明长度为1，访问第二个元素会报错
	a8 := [...]int{1, 2, 3}
	a8[0] = 111
	fmt.Printf("a8 %v\n", a8)
	fmt.Printf("a8 %v\n", a8[0])
	// 数组长度访问
	fmt.Printf("len(a8) %v\n", len(a8))
	// 遍历数组
	a9 := [3]int{1, 2, 3}
	for i := 0; i < len(a9); i++ {
		fmt.Printf("a9[i] %v\n", a9[i])
	}
	// for range遍历数组
	for ind, val := range a9 {
		fmt.Printf("%v %v\n", ind, val)
	}

	println()
	println("========================Array module console========================")
}
