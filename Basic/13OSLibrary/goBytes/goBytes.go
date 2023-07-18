package goBytes

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func GoBytes() {
	println("========================OS Library========================")
	println()

	// 类型转换
	println()
	typeTrans()

	// Contains
	println()
	containsFn()

	// Count
	println()
	countFn()

	// Repeat
	println()
	repeatFn()

	// Replace
	println()
	replaceFn()

	// Runes
	println()
	runesFn()

	// Join
	println()
	joinFn()

	// Reader
	println()
	readerFn()

	// Buffer
	println()
	bufferFn()

	println()
	println("========================OS Library========================")
}

func bufferFn() {
	// 方式一
	println()
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)

	n, _ := b.WriteString("hello")
	fmt.Printf("b.Bytes(): %v\n", string(b.Bytes()))
	fmt.Printf("b.Bytes()[0:n]: %v\n", string(b.Bytes()[0:n]))

	// 方式二
	// 读取固定的两个
	println()
	b1 := bytes.NewBufferString("hello")
	fmt.Printf("b1: %v\n", b1)
	mb := make([]byte, 2)
	read, _ := b1.Read(mb)
	fmt.Printf("read: %v\n", read)
	fmt.Printf("string(mb): %v\n", string(mb))
	// 循环读取所有
	println()
	b2 := bytes.NewBufferString("hello")
	fmt.Printf("b2: %v\n", b2)
	mb1 := make([]byte, 1)
	for {
		read1, err := b2.Read(mb1)
		if err == io.EOF {
			fmt.Printf("err: %v\n", err)
			break
		} else {
			fmt.Printf("read1: %v\n", read1)
			fmt.Printf("string(mb1): %v\n", string(mb1[0:read1]))
		}
	}

	// 方式三
	println()
	b3 := bytes.NewBuffer([]byte("hello"))
	fmt.Printf("b3: %v\n", b3)

}

func readerFn() {
	s := "123456789"
	// 创建Reader
	re := bytes.NewReader([]byte(s))
	// 返回未读取部分的长度
	fmt.Printf("re.Len(): %v\n", re.Len())
	// 返回底层数据总长度
	fmt.Printf("re.Size(): %v\n", re.Size())

	println("===============")

	buf := make([]byte, 2)
	for {
		// 读取数据
		n, err := re.Read(buf)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		} else {
			fmt.Printf("buf[:n]: %v\n", string(buf[:n]))
		}
	}

	println("===============")

	// 重置偏移量，因为上面的操作已经改变了读取位置等信息
	re.Seek(0, 0)

	for {
		b, err := re.ReadByte()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		} else {
			fmt.Printf("string(b): %v\n", string(b))
		}
	}

	println("===============")

	// 重置偏移量，因为上面的操作已经改变了读取位置等信息
	re.Seek(0, 0)

	off := int64(0)
	for {
		n, err := re.ReadAt(buf, off)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		} else {
			off += int64(n)
			fmt.Printf("string(buf[:n]): %v\n", string(buf[:n]))
		}
	}
}

func joinFn() {
	// 二维切片
	s1 := [][]byte{[]byte("你好"), []byte("世界")}
	// 连接符号声明
	sep1 := []byte(",")
	sep2 := []byte(" + ")
	fmt.Printf("以 \",\" 逗号为连接符: %v\n", string(bytes.Join(s1, sep1)))
	fmt.Printf("以 \" + \" 加号为连接符: %v\n", string(bytes.Join(s1, sep2)))

}

func runesFn() {
	s := []byte("你好世界")
	rs := bytes.Runes(s)
	fmt.Printf("转换前错误的汉字长度: %v\n", len(s))
	fmt.Printf("转换后正确的汉字长度: %v\n", len(rs))
}

func replaceFn() {
	s := []byte("hello world")
	oldVal := []byte("l")
	newVal := []byte("i")
	fmt.Printf("0表示不替换: %v\n", string(bytes.Replace(s, oldVal, newVal, 0)))
	fmt.Printf("1表示替换一次: %v\n", string(bytes.Replace(s, oldVal, newVal, 1)))
	fmt.Printf("2表示替换两次: %v\n", string(bytes.Replace(s, oldVal, newVal, 2)))
	fmt.Printf("-1表示替换所有不限次: %v\n", string(bytes.Replace(s, oldVal, newVal, -1)))
}

func repeatFn() {
	b := []byte("hi")
	fmt.Printf("string(bytes.Repeat(b,1)): %v\n", string(bytes.Repeat(b, 1)))
	fmt.Printf("string(bytes.Repeat(b,3)): %v\n", string(bytes.Repeat(b, 3)))
}

func countFn() {
	s := []byte("helloooooooo")
	b1 := []byte("h")
	b2 := []byte("l")
	b3 := []byte("o")
	fmt.Printf("bytes.COunt(s,b1): %v\n", bytes.Count(s, b1))
	fmt.Printf("bytes.COunt(s,b2): %v\n", bytes.Count(s, b2))
	fmt.Printf("bytes.COunt(s,b3): %v\n", bytes.Count(s, b3))
}

func containsFn() {
	s := "leechenze.com"
	b := []byte(s)

	b1 := []byte("leechenze")
	b2 := []byte("LeeChenZe")

	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)
	b4 := bytes.Contains(b, b2)
	fmt.Printf("b4: %v\n", b4)
	println()
	fmt.Printf("strings contains: %v\n", strings.Contains("hello world", "hello"))
}

func typeTrans() {
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)
	println(b)
	println(i)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)

}
