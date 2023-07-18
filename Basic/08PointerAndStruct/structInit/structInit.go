package structInit

import "fmt"

type Person struct {
	id, age     int
	name, email string
}

// 结构体作为函数参数（值传递，并不会改变外部Person的值。）
func changePerson(person Person) {
	person.id = 105
	person.name = "nixon"
	fmt.Printf("nixon, %v\n", person)
}

// 结构体作为函数参数（引用传递/指针传递，将会改变外部Person的值。）
func changeperson2(person *Person) {
	person.id = 107
	person.name = "nixon"
	fmt.Printf("nixon, %v\n", person)
}

func StructInit() {
	println("========================Pointer And Struct========================")
	println()

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

	lincoln := Person{
		103, 30, "lincoln", "lincoln@gmail.com",
	}
	var p_person *Person
	p_person = &lincoln
	fmt.Printf("p_person, %p\n", &lincoln)
	fmt.Printf("*p_person, %v\n", *p_person)

	// 使用new关键字创建结构体指针
	trump := new(Person)
	trump.id = 104
	trump.name = "trump"
	trump.age = 80
	fmt.Printf("trump, %v, %p\n", trump, trump)
	fmt.Printf("trump, %v, \n", *trump)

	// 结构体作为函数参数
	println()
	println("结构体作为函数参数 - 值传递，并不会改变外部Person的值。")
	douglas := Person{
		id:   106,
		name: "douglas",
	}
	fmt.Printf("douglas, %v\n", douglas)
	println("==============================")
	changePerson(douglas)
	fmt.Printf("douglas, %v\n", douglas)

	println()
	println("结构体作为函数参数 - 引用传递/指针传递，将会改变外部Person的值。")
	var p_douglas = &douglas
	fmt.Printf("douglas, %v\n", douglas)
	println("==============================")
	changeperson2(p_douglas)
	fmt.Printf("douglas, %v\n", douglas)

	// 结构体嵌套
	println()
	println("结构体嵌套")
	type Dog struct {
		name, color string
		age         int
	}
	type Owner struct {
		dog  Dog
		name string
		age  int
	}

	dog := Dog{
		name:  "bagong",
		age:   8,
		color: "yellow",
	}
	owner := Owner{
		name: "roosevelt",
		age:  50,
		dog:  dog,
	}
	fmt.Printf("owner, %v\n", owner)
	fmt.Printf("owner.dog.name, %v\n", owner.dog.name)

	println()
	println("========================Pointer And Struct========================")

}
