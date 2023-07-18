package gomap

import "fmt"

func GoMap() {
	println("========================Map start========================")
	println()

	// 声明map
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1 %v\n", m1)

	// 初始化map
	println()
	println("初始化map方式一")
	m2 := map[string]string{"name": "leechenze", "age": "20", "email": "leeczyc@gmail.com"}
	fmt.Printf("m2 %v\n", m2)
	println()
	println("初始化map方式二")
	m3 := make(map[string]string)
	m3["name"] = "leechenze"
	m3["age"] = "20"
	m3["email"] = "leeczyc@gmail.com"
	fmt.Printf("m3 %v\n", m3)

	// 访问map
	println()
	println("访问map")
	name := m3["name"]
	fmt.Printf("name %v\n", name)

	// 判断某个Key是否存在
	println()
	println("判断某个Key是否存在")
	val, ok := m3["name"]
	fmt.Printf("val %v\n", val)
	fmt.Printf("ok %v\n", ok)

	// map遍历
	println()
	println("range 遍历map")
	for key, val := range m3 {
		fmt.Printf("key %v\n", key)
		fmt.Printf("val %v\n", val)
	}

	println()
	println("========================Map end========================")
}
