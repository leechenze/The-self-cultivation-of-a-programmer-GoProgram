package basicPointer

import "fmt"

func BasicPointer() {
	println("========================Basic Pointer========================")
	println()

	// 声明指针
	var ip *int
	fmt.Printf("ip %v\n", ip)

	// 指针赋值
	var i int = 100
	ip = &i
	fmt.Printf("ip %v\n", ip)

	// 指针取值
	var ipVal = *ip
	fmt.Printf("ipVal %v\n", ipVal)

	// 指针字符串
	println()
	println("指针字符串")
	var sp *string
	var str string = "hello"
	sp = &str
	fmt.Printf("sp %v\n", sp)

	// 指针布尔
	var bp *bool
	var flag bool = false
	bp = &flag
	fmt.Printf("bp %v\n", bp)
	fmt.Printf("*bp %v\n", *bp)

	println()
	println("========================Basic Pointer========================")
}
