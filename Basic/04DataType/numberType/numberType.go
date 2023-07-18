package numberType

import (
	"fmt"
	"math"
	"unsafe"
)

func NumberType() {
	println("========================Number Type start========================")
	println()

	// 整数
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	// 无符号整数
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)

	fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
	fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)
	fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)
	fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

	println()
	// 以二进制，八进制或十六进制浮点数的格式定义数字
	var a int = 10
	// 十进制（decimalism）
	fmt.Printf("%d \n", a)
	// 二进制（binary）
	fmt.Printf("%b \n", a)
	// 八进制（Octal）以0开头。
	var b int = 077
	fmt.Printf("%o \n", b)
	// 十六进制（hexadecimal）以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	println()
	// 浮点型
	fmt.Printf("%f\n", math.Pi)
	// .2代表保留2位小数
	fmt.Printf("%.2f\n", math.Pi)

	println()
	println("========================Number Type start========================")
}
