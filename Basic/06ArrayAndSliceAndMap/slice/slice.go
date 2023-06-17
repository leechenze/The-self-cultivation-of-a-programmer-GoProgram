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

	// 切片切割，功能就是JS数组的slice方法。
	println()
	var a5 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	a6 := a5[:3]
	a66 := a5[0:3]
	fmt.Printf("a6 %v\n", a6)
	fmt.Printf("a66 %v\n", a66)
	a7 := a5[3:]
	fmt.Printf("a7 %v\n", a7)
	a8 := a5[2:5]
	fmt.Printf("a8 %v\n", a8)
	a9 := a5[:]
	fmt.Printf("a9 %v\n", a9)

	// 切片的CRUD操作
	println()
	println("add")
	var a10 = []int{}
	fmt.Printf("a10 %v\n", a10)
	a10 = append(a10, 100)
	a10 = append(a10, 200)
	a10 = append(a10, 300)
	fmt.Printf("a10 %v\n", a10)
	println()
	println("del")
	// 删除下标为2的切片（3）
	a11 := []int{1, 2, 3, 4}
	fmt.Printf("a11[:2] %v\n", a11[:2])
	fmt.Printf("a11[3:] %v\n", a11[3:])
	a11 = append(a11[:2], a11[3:]...)
	fmt.Printf("a11 %v\n", a11)

	println()
	println("update")
	a12 := []int{1, 2, 3, 4}
	a12[1] = 100
	fmt.Printf("a12 %v\n", a12)

	// 获取切片下标
	println()
	println("获取索引下标")
	a13 := []string{"hello", "world", "leechenze", "niubility"}
	for ind, val := range a13 {
		if val == "niubility" {
			fmt.Printf("%v\n", ind)
		}
	}

	// 切片拷贝，如果直接将一个切片赋值给另一个切片，那么这两个切片指向同一个地址，修改一个切片后另一个切片也会被修改。
	// 这个问题可以用拷贝函数来处理，拷贝值要用make创建
	println()
	println("切片拷贝（深拷贝）")
	a14 := []int{1, 2, 3, 4, 5}
	a15 := make([]int, len(a14))
	copy(a15, a14)
	fmt.Printf("a14修改前 %v\n", a14)
	fmt.Printf("a15修改前 %v\n", a15)
	a14[0] = 100
	fmt.Printf("a14修改后 %v\n", a14)
	fmt.Printf("a15修改后 %v\n", a15)

	println()
	println("========================Slice module console========================")
}
