package formatOutput

import "fmt"

type WebSite struct {
	Name string
}

func FormatOutput() {
	println("========================Format Output start========================")
	println()

	println()
	println("普通占位符")
	site := WebSite{Name: "leechenze"}
	// %v:相应值的默认格式，相当于value
	fmt.Printf("%v\n", site)
	// %#v:相应值的默认格式，包括结构类型也会输出
	fmt.Printf("%#v\n", site)
	// %T:输出类型
	fmt.Printf("%T\n", site)
	// %%:输出%的字面量
	fmt.Printf("%%\n", site)

	println()
	println("布尔占位符")
	flag := true
	// %t:输出布尔类型
	fmt.Printf("%t\n", flag)

	println()
	println("整数占位符")
	// %b: 输出二进制
	fmt.Printf("%b\n", 5)
	// %c: 相应Unicode码所表示的字符
	fmt.Printf("%c\n", 0x4E2D)
	// %d: 十进制表示
	fmt.Printf("%d\n", 0x12)
	// %o：八进制表示
	fmt.Printf("%o\n", 10)
	// %q：单引号围绕的字符字面值，由go语法安全的转义
	fmt.Printf("%q\n", 0x4E2D)
	// %x：十六进制表示，字母形式为小写 a-f
	fmt.Printf("%x\n", 13)
	// %X：十六进制表示，字母形式为大写 A-F
	fmt.Printf("%X\n", 13)
	// %U：Unicode格式：U+1234，等同于 "U+%04X"
	fmt.Printf("%U\n", 0x4E2D)

	println()
	println("字符串与字节切片")
	// %s：输出字符串表示(string类型或[]byte)
	fmt.Printf("%s\n", []byte("leechenze"))
	fmt.Printf("%v\n", []byte("leechenze"))

	println()
	println("指针")
	// %p: 十六进制表示，前缀0x的指针地址输出
	fmt.Printf("%p", &site)

	println()
	println("========================Format Output start========================")

}
