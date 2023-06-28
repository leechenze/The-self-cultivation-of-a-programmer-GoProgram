package goMath

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GoMath() {
	println("========================OS Library========================")
	println()

	// PI，%.2f表示留两位小数
	fmt.Printf("math.Pi: %.2f\n", math.Pi)

	/** 常用函数 */
	println()
	fmt.Printf("取绝对值: %f\n", math.Abs(-3.14))
	fmt.Printf("取x的y次方: %f\n", math.Pow(2, 4))
	fmt.Printf("取x的开平方: %f\n", math.Sqrt(64))
	fmt.Printf("取x的开立方: %f\n", math.Cbrt(27))
	fmt.Printf("向上取整: %f\n", math.Ceil(3.14))
	fmt.Printf("向下取整: %f\n", math.Floor(3.14))
	fmt.Printf("四舍五入：%f\n", math.Round(3.14))
	fmt.Printf("取余数: %f\n", math.Mod(10, 3))
	integer, decimal := math.Modf(3.1415926535)
	fmt.Printf("整数部分为: %f，小数部分为：%f\n", integer, decimal)

	/** 随机数 */
	println()
	// 如果Seed（种子）是一个固定值，那么每次随机数执行的结果都是相同的。
	// rand.Seed(1)
	// 所以将Seed设置为一个时间戳，保证每次的随机数都不相同
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		fmt.Printf("随机数: %v\n", rand.Int())
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Printf("100以内的随机数: %v\n", rand.Intn(100))
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Printf("随机小数（0-1）: %v\n", rand.Float64())
	}

	println()
	println("========================OS Library========================")
}
