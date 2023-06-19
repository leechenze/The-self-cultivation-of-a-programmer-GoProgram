package arrayPointer

import "fmt"

func ArrayPointer() {
	println("========================Pointer And Struct========================")
	println()

	// 数组指针声明
	var ap [3]*int
	arr := [3]int{1, 2, 3}
	fmt.Printf("ap %v\n", ap)

	// 数组指针赋值
	// 通过for循环对 arr 数组的每一个元素进行赋值
	for i := 0; i < len(arr); i++ {
		ap[i] = &arr[i]
	}
	fmt.Printf("ap %v\n", ap)

	// 数组指针取值
	for n := 0; n < len(arr); n++ {
		fmt.Printf("*ap[n] %v\n", *ap[n])
	}

	println()
	println("========================Pointer And Struct========================")

}
