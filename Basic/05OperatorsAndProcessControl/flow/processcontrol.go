package flow

import "fmt"

// 判断一个数是奇数还是偶数
func foo() {
	var num int
	// Scanln用于接收用户的输入, &是取地址符号，用于取num变量的地址。
	fmt.Scanln(&num)
	println(&num)
	println("请输入一个数字")
	if num%2 == 0 {
		println("偶数")
	} else {
		println("奇数")
	}
}

// goto关键字演示
func gotoTest() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i >= 2 && j >= 2 {
				goto END
			}
			println(i, j)
		}
	}
END:
	println("goto结束循环，跳到了此处")
}

func fallthroughTest() {
	switch 100 {
	case 100:
		println(100)
		fallthrough
	case 200:
		println(200)
	case 300:
		println(300)
	default:
		println("other")
	}
}

func Flow() {
	println("========================Operators start========================")
	println()

	a := 100

	println()
	println("if else statement")
	if a >= 20 {
		println(true)
	} else {
		println(false)
	}
	// 初始变量可以声明在布尔表达式里面
	if age := 20; age > 18 {
		println("成年了")
	} else {
		println("未成年")
	}

	println()
	println("switch statement")
	switch a {
	case 100:
		println("100")
	default:
		println("default")
	}
	// switch 多条件匹配
	day := 3
	switch day {
	case 1, 2, 3, 4, 5:
		println("工作日")
	case 6, 7:
		println("休息日")
	}
	println()
	// fallthrough关键字演示
	fallthroughTest()

	println()
	println("for Loop")
	for i := 0; i < 10; i++ {
		println(i)
	}
	println()
	println("for range Loop")
	x := [...]string{"lincoln", "trump", "obama"}
	for ind, val := range x {
		println(ind, val)
	}

	println()
	println("判断一个数是奇数还是偶数")
	// foo()

	println()
	println("多重判断")
	score := 80
	if score >= 60 && score <= 70 {
		println("C")
	} else if score >= 70 && score <= 90 {
		println("B")
	} else if score >= 90 {
		println("A")
	}

	println()
	println("goto关键字演示")
	gotoTest()

	println()
	println("========================Operators end========================")
}
