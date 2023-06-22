package constructor

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name 不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age 不能小于0")
	}
	return &Person{name: "tom", age: 20}, nil
}

func Constructor() {
	println("========================constructor========================")
	println()

	person, err := NewPerson("trump", 40)
	if err == nil {
		fmt.Printf("person %v\n", person)
	} else {
		fmt.Printf("err %v\n", err)
	}

	println()
	println("========================constructor========================")
}
