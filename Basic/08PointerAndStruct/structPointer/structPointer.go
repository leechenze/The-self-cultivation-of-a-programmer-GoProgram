package structPointer

import "fmt"

// 定义结构体
type Person struct {
	id    int
	name  string
	age   int
	email string
}
type Customer struct {
	id, age     int
	name, email string
}

func StructPointer() {
	println("========================Pointer And Struct========================")
	println()

	// 类型定义与类型别名
	println()
	println("类型定义")
	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("i, %v, %T\n", i, i)
	println()
	println("类型别名")
	type MyInt1 = int
	var i1 MyInt1
	i1 = 200
	fmt.Printf("i1, %v, %T\n", i1, i1)

	// 定义结构体
	println()
	println("结构体")
	var tom Person
	fmt.Printf("tom %v\n", tom)
	var kate Customer
	fmt.Printf("kate %v\n", kate)

	// 结构体赋值
	tom.id = 101
	tom.name = "tom"
	tom.age = 20
	tom.email = "tom@gmail.com"
	fmt.Printf("tom %v\n", tom)

	// 结构体访问
	fmt.Printf("tom.name, %v\n", tom.name)

	// 匿名结构体
	var clinton struct {
		id, age     int
		name, email string
	}
	clinton.id = 102
	clinton.name = "clinton"
	clinton.age = 30
	clinton.email = "clinton@gmail.com"
	fmt.Printf("clinton, %v\n", clinton)
	fmt.Printf("clinton.email, %v\n", clinton.email)

	// 结构体指针部分请看: structInit/structInit.go

	println()
	println("========================Pointer And Struct========================")

}
