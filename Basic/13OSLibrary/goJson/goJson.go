package goJson

import (
	"encoding/json"
	"fmt"
	"os"
)

func GoJson() {
	println("========================OS Library========================")
	println()

	// 基础json转换
	println()
	basicJson()

	// 嵌套结构体转换
	println()
	nestedJson()

	// Decoder结合IO读取Json文件
	println()
	decoderFn()

	// Encoder结合IO写入Json文件
	println()
	encoderFn()

	println()
	println("========================OS Library========================")

}

func encoderFn() {
	person := Person1{
		Name:   "clinton",
		Age:    50,
		Email:  "tom@gmail.com",
		Parent: []string{"big Clinton", "big Hillary"},
	}

	file, _ := os.OpenFile("operate/write.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(person)
}

func decoderFn() {
	file, _ := os.Open("operate/read.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	// interface{} 其实就是 any 任意类型
	var val map[string]interface{}
	decoder.Decode(&val)
	fmt.Printf("val: %v\n", val)
}

func nestedJson() {
	byteSlice1 := []byte("{\"Name\":\"tom\",\"Age\":20,\"Email\":\"tom@gmail.com\", \"Parents\": [\"big tom\", \"big kite\"]}")
	var person2 interface{}
	// 注意Unmarshal第二个参数应该是指针类型，需要&取地址
	json.Unmarshal(byteSlice1, &person2)
	fmt.Printf("person2: %v\n", person2)
	// 将person2 这个interface类型强转为map类型。
	for key, val := range person2.(map[string]interface{}) {
		fmt.Printf("key:val = %v %v \n", key, val)
	}
	person3 := Person1{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}
	jsonPerson1, _ := json.Marshal(person3)
	fmt.Printf("jsonPerson1: %v\n", string(jsonPerson1))
}

func basicJson() {
	person := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	jsonPerson, _ := json.Marshal(person)
	fmt.Printf("jsonPerson: %v\n", string(jsonPerson))

	byteSlice := []byte("{\"Name\":\"tom\",\"Age\":20,\"Email\":\"tom@gmail.com\"}\n")
	var person1 Person
	// 注意Unmarshal第二个参数应该是指针类型，需要&取地址
	json.Unmarshal(byteSlice, &person1)
	fmt.Printf("person1: %v\n", person1)
}

type Person struct {
	Name  string
	Age   int
	Email string
}

type Person1 struct {
	Name   string
	Age    int
	Email  string
	Parent []string
}
