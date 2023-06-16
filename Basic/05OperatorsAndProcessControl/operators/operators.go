package operators

import "fmt"

func Operators() {
	println("========================Operators start========================")
	println()

	println("算数运算符")
	a := 100
	b := 20
	println("(a + b)", (a + b))
	println("(a - b)", (a - b))
	println("(a * b)", (a * b))
	println("(a / b)", (a / b))
	println("(a % b)", (a % b))

	println()
	println("自增自减")
	c := 1
	c++
	println(c)

	println()
	println("关系运算符")
	println("(a > b)", (a > b))
	println("(a < b)", (a < b))
	println("(a >= b)", (a >= b))
	println("(a <= b)", (a <= b))
	println("(a == b)", (a == b))
	println("(a != b)", (a != b))

	println()
	println("逻辑运算符")
	d := true
	e := false
	println((d && e))
	println((d || e))
	println((!d))
	println((!e))

	println()
	println("逻辑运算符")
	f := 4 // 二进制 100
	fmt.Printf("a: %b\n", f)
	g := 8 // 二进制 1000
	fmt.Printf("a: %b\n", g)
	fmt.Printf("(f & g): %v %b \n", (f & g), (f & g))
	fmt.Printf("(f | g): %v %b \n", (f | g), (f | g))
	fmt.Printf("(f ^ g): %v %b \n", (f ^ g), (f ^ g))
	fmt.Printf("(f << g): %v %b \n", (f << g), (f << g))
	fmt.Printf("(f >> g): %v %b \n", (f >> g), (f >> g))

	println()
	println("赋值运算符")
	h := 100
	h = 200
	println("h", h)
	h += 1
	println("h += 1", h)
	h -= 1
	println("h -= 1", h)
	h *= 1
	println("h *= 1", h)
	h /= 1
	println("h /= 1", h)

	println()
	println("========================Operators end========================")

}
