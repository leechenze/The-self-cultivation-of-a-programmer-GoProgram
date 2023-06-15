package stringType

import (
	"bytes"
	"fmt"
	"strings"
)

func StringType() {
	println("========================String Type start========================")
	println()

	var str1 string = "hello"
	var htm1 string = `
		<html>
			<head><title>hello golang</title></head>
		</html>
	`

	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("htm1: %v\n", htm1)

	// 字符串连接
	s1 := "tom"
	s2 := "20"
	msg := s1 + s2
	fmt.Printf("msg: %v \n", msg)

	// buffer.WriteString()
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	println()
	println("字符串切片操作")
	// 字符串切片操作
	str2 := "Hello World"
	n := 3
	m := 5
	// 获取字符串索引位置为n的字符的原始字节
	fmt.Printf("%v\n", str2[n])
	// 获取字符串索引位置为n的字符
	fmt.Printf("%c\n", str2[n])
	// 截取得字符串索引位置为 n 到 m-1 的字符串
	fmt.Println(str2[n:m])
	// 截取得字符串索引位置为 n 到 len(s)-1 的字符串
	fmt.Println(str2[n:])
	// 截取得字符串索引位置为 0 到 m-1 的字符串
	fmt.Println(str2[:m])

	println()
	println("字符串函数")
	// 字符串长度
	println(len(str2))
	// 分隔为数组
	fmt.Printf("%v\n", strings.Split(str2, " "))
	// 判断是否包含
	fmt.Printf("%v\n", strings.Contains(str2, "Hello"))
	// 判断是否是以 hello 为前缀的
	fmt.Printf("%v\n", strings.HasPrefix(str2, "Hello"))
	// 判断是否是以 world 为后缀的
	fmt.Printf("%v\n", strings.HasSuffix(str2, "World"))
	// 转换大写
	fmt.Printf("%v\n", strings.ToUpper(str2))
	// 转换小写
	fmt.Printf("%v\n", strings.ToLower(str2))
	// 返回指定字符的索引位置
	fmt.Printf("%v\n", strings.Index(str2, "ll"))
	// 返回指定字符的索引位置
	fmt.Printf("%v\n", strings.LastIndex(str2, "ll"))

	println()
	println("========================String Type start========================")
}
