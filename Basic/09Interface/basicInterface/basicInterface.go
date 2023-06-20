package basicInterface

import "fmt"

// 定义接口
type USB interface {
	read()
	write()
}

// 定义电脑的结构体来实现USB接口方法
type Computer struct {
	name string
}

// 定义手机的结构体来实现USB接口方法
type Mobile struct {
	brand string
}

// Mobile结构体的方法，实现USB接口read方法
func (m Mobile) read() {
	fmt.Printf("m.brand %v\n", m.brand)
	fmt.Println("mobile read...")
}
func (m Mobile) write() {
	fmt.Printf("m.brand %v\n", m.brand)
	fmt.Println("mobile write...")
}

// Computer结构体的方法，实现USB接口read方法
func (c Computer) read() {
	fmt.Printf("c.name %v\n", c.name)
	fmt.Println("computer read...")
}
func (c Computer) write() {
	fmt.Printf("c.name %v\n", c.name)
	fmt.Println("computer write...")
}

// =============================================

// 接口的值类型接收者和指针类型接收者
type Peter interface {
	eat(string) string
}
type Dog struct {
	name string
}
type Cat struct {
	name string
}

func (dog Dog) eat(food string) string {
	dog.name = "豆豆"
	fmt.Printf("food, %v\n", food)
	return dog.name + "吃" + food
}
func (cat *Cat) eat(food string) string {
	cat.name = "欢欢"
	fmt.Printf("food, %v\n", food)
	return cat.name + "吃" + food
}

func BasicInterface() {
	println("========================Interface========================")
	println()

	c := Computer{
		name: "lenovo",
	}

	c.read()
	c.write()

	println()
	m := Mobile{
		brand: "apple",
	}

	m.read()
	m.write()

	// 接口的值类型接收者和指针类型接收者
	println()
	println("接口的值类型接收者和指针类型接收者")
	dog := Dog{
		name: "徐坤",
	}
	dog.eat("火鸡")
	fmt.Printf("dog, %v\n", dog)
	println("============================")
	cat := &Cat{
		name: "徐坤",
	}
	cat.eat("鱼干")
	fmt.Printf("cat, %v\n", cat)

	// 接口和类型的关系
	println()
	println("接口和类型的关系")
	// TODO

	println()
	println("========================Interface========================")
}
