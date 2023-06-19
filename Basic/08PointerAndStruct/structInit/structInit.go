package structInit

import "fmt"

func StructInit() {
	println("========================Pointer And Struct========================")
	println()

	type Person struct {
		id, age     int
		name, email string
	}

	// 结构体初始化（键值对初始化）
	var tom Person
	tom = Person{
		id:    101,
		name:  "tom",
		age:   20,
		email: "tom@gmail.com",
	}
	fmt.Printf("tom, %v\n", tom)

	// 通过默认的声明顺序进行初始化（列表初始化）
	var clinton = Person{
		102, 20, "clinton", "clinton@gmail.com",
	}
	fmt.Printf("clinton, %v\n", clinton)

	// 结构体指针
	println()
	println("结构体指针")

	println()
	println("========================Pointer And Struct========================")

}
