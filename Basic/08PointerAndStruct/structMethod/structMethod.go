package structMethod

import "fmt"

type Person struct {
	name string
}

// （per Person）接收者reciver，以示eat方法是属于Person结构体的。
func (per Person) eat() {
	fmt.Printf("eat..., %v\n", per.name)
}
func (per Person) sleep() {
	fmt.Printf("sleep..., %v\n", per.name)
}

// 结构体方法的值类型和指针类型（参数实现）
func changePerson1(person Person) {
	person.name = "Roosevelt"
}
func changePerson2(person *Person) {
	person.name = "Benjamin"
}

// 结构体方法的值类型和指针类型（接收者实现）
func (person Person) changePerson3() {
	person.name = "Roosevelt"
}
func (person *Person) changePerson4() {
	person.name = "Benjamin"
}

func StructMethod() {
	println("========================Pointer And Struct========================")
	println()

	// golang结构体中的方法
	person := Person{
		name: "tom",
	}
	person.eat()
	person.sleep()

	println()
	println("结构体方法的值类型和指针类型（接收者实现）")
	println()
	// 结构体方法的值类型和指针类型（参数实现）
	person1 := Person{
		name: "person1",
	}
	person2 := &Person{
		name: "person2",
	}
	// changePerson1执行后，person1并不会改变，因为是值传递。
	fmt.Printf("person1, %v\n", person1)
	changePerson1(person1)
	fmt.Printf("person1, %v\n", person1)
	println("================================")
	// changePerson2执行后，person2就改变了，因为是指针传递/引用传递。
	fmt.Printf("person2, %v\n", person2)
	changePerson2(person2)
	fmt.Printf("person2, %v\n", person2)

	println()
	println("结构体方法的值类型和指针类型（接收者实现）")
	println()

	// 结构体方法的值类型和指针类型（接收者实现）
	person3 := Person{
		name: "person3",
	}
	person4 := &Person{
		name: "person4",
	}
	// changePerson3执行后，person3并不会改变，因为是值传递。
	fmt.Printf("person3, %v\n", person3)
	person3.changePerson3()
	fmt.Printf("person3, %v\n", person3)
	println("================================")
	// changePerson4执行后，person4就改变了，因为是指针传递/引用传递。
	fmt.Printf("person4, %v\n", person4)
	person4.changePerson4()
	fmt.Printf("person4, %v\n", person4)

	println()
	println("========================Pointer And Struct========================")
}
