package basicInherit

import "fmt"

type Animal struct {
	name string
	age  int
}

func (animal Animal) eat() {
	println("animal eat")
}
func (animal Animal) sleep() {
	println("animal sleep")
}

type Dog struct {
	animal         Animal
	characteristic string
}
type Cat struct {
	Animal
	characteristic string
}

func BasicInherit() {
	println("========================Inherit========================")
	println()

	dog := Dog{
		animal:         Animal{name: "carter", age: 6},
		characteristic: "loyalty",
	}
	dog.animal.eat()
	dog.animal.sleep()
	fmt.Printf("dog.characteristic %v\n", dog.characteristic)
	fmt.Printf("dog.animal.name %v\n", dog.animal.name)
	fmt.Printf("dog.animal.age %v\n", dog.animal.age)
	// 或者还可以如下按着顺序则无需写属性
	println()
	cat := Cat{
		Animal{name: "ethan", age: 8},
		"lovely",
	}
	cat.eat()
	cat.sleep()
	fmt.Printf("cat.characteristic %v\n", cat.characteristic)
	fmt.Printf("cat.name %v\n", cat.name)
	fmt.Printf("cat.age %v\n", cat.age)
	println()
	println("========================Inherit========================")
}
