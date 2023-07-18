package goIO

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/** ioutil */
// ReadAll演示
func readAll() {
	// 文件
	// byt := strings.NewReader("hello world")
	//
	dir, _ := os.Getwd()
	os.Chdir(dir + "/operate")
	byt, _ := os.Open("a1.txt")
	defer byt.Close()
	all, err := ioutil.ReadAll(byt)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("string(all): %v\n", string(all))
	}
}

// ReadDir演示
func readDir() {
	// .表示当前目录，因为Chdir到了operate目录下，结果是operate目录下的内容。
	dir, _ := ioutil.ReadDir(".")
	for _, val := range dir {
		fmt.Printf("val.Name(): %v\n", val.Name())
	}
}

// readFile
func readFile() {
	file, _ := ioutil.ReadFile("a1.txt")
	fmt.Printf("string(file): %v\n", string(file))
}

// writeFile演示
func writeFile() {
	err := ioutil.WriteFile("a5.txt", []byte("hello world"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("a5.txt文件写入成功")
	}
}

/** bufio */
// NewReader演示
func newReader() {
	// byt := strings.NewReader("hello world")
	byt, _ := os.Open("a5.txt")
	defer byt.Close()
	readerObj := bufio.NewReader(byt)
	readstr, _ := readerObj.ReadString(' ')
	fmt.Printf("readstr: %v\n", readstr)
}

// read演示
func read() {
	s := strings.NewReader("ABCDEFGHIGKLMNOPQRSTUVWXYZ")
	byteRead := bufio.NewReader(s)
	p := make([]byte, 10)
	for {
		n, err := byteRead.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p[0:n]): %v\n", string(p[0:n]))
		}
	}
}

// readByte和unreadByte演示
func readByte() {
	str := strings.NewReader("ABCDEFGHIGKLMNOPQRSTUVWXYZ")
	br := bufio.NewReader(str)

	b, _ := br.ReadByte()
	fmt.Printf("string(b): %v\n", string(b))

	b, _ = br.ReadByte()
	fmt.Printf("string(b): %v\n", string(b))

	// 吐出最近一次读取操作的最后一个直接
	br.UnreadByte()

	b, _ = br.ReadByte()
	fmt.Printf("string(b): %v\n", string(b))
}

// readRune和unreadRune演示，如果是中文，韩文等字符需要用Rune进行操作。
func readRune() {
	str := strings.NewReader("你好，世界！")
	br := bufio.NewReader(str)

	b, size, _ := br.ReadRune()
	fmt.Printf("string(b): %v,%v \n", string(b), size)

	b, size, _ = br.ReadRune()
	fmt.Printf("string(b): %v,%v \n", string(b), size)

	// 吐出最近一次读取操作的最后一个直接
	br.UnreadRune()

	b, size, _ = br.ReadRune()
	fmt.Printf("string(b): %v,%v \n", string(b), size)
}

// readLine演示
func readLine() {
	str := strings.NewReader("ABCD\nEFG\nHIGK\nLMN\n")
	br := bufio.NewReader(str)

	line, prefix, _ := br.ReadLine()
	fmt.Printf("line,prefix: %v %v \n", string(line), prefix)

	line, prefix, _ = br.ReadLine()
	fmt.Printf("line,prefix: %v %v \n", string(line), prefix)

	line, prefix, _ = br.ReadLine()
	fmt.Printf("line,prefix: %v %v \n", string(line), prefix)

	line, prefix, _ = br.ReadLine()
	fmt.Printf("line,prefix: %v %v \n", string(line), prefix)
}

// readSlice演示
func readSlice() {
	str := strings.NewReader("ABC,DEFG,HIGK,LMN")
	br := bufio.NewReader(str)

	word, _ := br.ReadSlice(',')
	fmt.Printf("word: %v\n", string(word))

	word, _ = br.ReadSlice(',')
	fmt.Printf("word: %v\n", string(word))

	word, _ = br.ReadSlice(',')
	fmt.Printf("word: %v\n", string(word))

	word, _ = br.ReadSlice(',')
	fmt.Printf("word: %v\n", string(word))
}

// readBytes演示
func readBytes() {
	str := strings.NewReader("ABC DEFG HIGK LMN")
	br := bufio.NewReader(str)

	word, _ := br.ReadBytes(' ')
	fmt.Printf("word: %v\n", string(word))

	word, _ = br.ReadBytes(' ')
	fmt.Printf("word: %v\n", string(word))

	word, _ = br.ReadBytes(' ')
	fmt.Printf("word: %v\n", string(word))

	word, _ = br.ReadBytes(' ')
	fmt.Printf("word: %v\n", string(word))
}

// readString演示
func readString() {
	str := strings.NewReader("ABC DEFG HIGK LMN")
	br := bufio.NewReader(str)

	word, _ := br.ReadString(' ')
	fmt.Printf("readString word: %v\n", string(word))

	word, _ = br.ReadString(' ')
	fmt.Printf("readString word: %v\n", string(word))

	word, _ = br.ReadString(' ')
	fmt.Printf("readString word: %v\n", string(word))

	word, _ = br.ReadString(' ')
	fmt.Printf("readString word: %v\n", string(word))
}

// writeTo演示
func writeTo() {
	str := strings.NewReader("ABCDEFGHIGKLMN")
	br := bufio.NewReader(str)
	// 创建缓冲区
	// byt := bytes.NewBuffer(make([]byte, 10))
	// 也可以打开一个文件作为写入对象（缓冲区）
	byt, err := os.OpenFile("a6.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer byt.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("a6.txt文件写入成功")
	}

	// 写入缓冲区
	// br.WriteTo(byt)
	// 写入文件
	br.WriteTo(byt)
	fmt.Printf("byt: %v\n", byt)
}

// newWriter演示
func newWriter() {
	// 加入openfile的不是一个文件，而是一个视频，那么bufio写的就会更高效。因为具有缓冲。
	file, err := os.OpenFile("a7.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	bw := bufio.NewWriter(file)
	defer file.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("a7.txt文件写入成功")
	}
	bw.WriteString("hello world")
	// 要刷新缓冲区，否则会一直在缓冲区，写不进去
	bw.Flush()
}

// reset演示
func reset() {
	b1 := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b1)
	bw.WriteString("123")
	b2 := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(b2)
	bw.WriteString("456")
	bw.Flush()
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
}

// writeByte和writeRune演示
func writeByte() {
	b1 := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b1)
	// 写入缓存 byte等同于 int8
	bw.WriteByte('H')
	bw.WriteByte('e')
	bw.WriteByte('l')
	bw.WriteByte('l')
	bw.WriteByte('o')
	bw.WriteByte(' ')
	// 写入缓存 rune等同于 int32
	bw.WriteRune('世')
	bw.WriteRune('界')
	bw.WriteRune('！')
	// 刷新缓存
	bw.Flush()
	fmt.Printf("b1: %v\n", b1)
}

// readFrom演示
func readFrom() {
	b1 := bytes.NewBuffer(make([]byte, 0))
	sr := strings.NewReader("Hello 世界！")
	bw := bufio.NewWriter(b1)
	bw.ReadFrom(sr)
	fmt.Printf("b1: %v\n", b1)
}

// newReadWriter演示
func newReadWriter() {
	b1 := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b1)
	br := bufio.NewReader(strings.NewReader("123"))
	brw := bufio.NewReadWriter(br, bw)
	brwStr, _ := brw.ReadString('\n')
	fmt.Printf("brwStr: %v\n", string(brwStr))
	brw.WriteString("asdf")
	brw.Flush()
	fmt.Printf("b1: %v\n", b1)

}

// newScanner演示
func newScanner() {
	sr := strings.NewReader("ABC DEF GHI JKL 梵蒂冈")
	bs := bufio.NewScanner(sr)
	// 以空格进行分隔
	// bs.Split(bufio.ScanWords)
	// 以字节进行分隔
	// bs.Split(bufio.ScanBytes)
	// 以特殊字符进行分隔
	bs.Split(bufio.ScanRunes)
	// Scan方法返回值为true或false，有扫描到结果返回true，否则返回false。
	for bs.Scan() {
		fmt.Printf("bs.Text(): %v\n", bs.Text())
	}
}

func GoIO() {
	println("========================OS Library========================")
	println()

	/** ioutil */
	// readAll演示
	readAll()

	// readDir演示
	readDir()

	// readFile演示
	readFile()

	// writeFile演示
	writeFile()

	/** bufio */
	// NewReader演示
	newReader()

	// read演示
	read()

	// readByte和unreadByte演示
	readByte()

	// readRune和unreadRune演示
	readRune()

	// readLine演示
	readLine()

	// readSlice演示
	readSlice()

	// readBytes演示
	readBytes()

	// readString演示
	readString()

	// writeTo演示
	writeTo()

	// newWriter演示
	newWriter()

	// reset演示
	reset()

	// writeByte和writeRune演示
	writeByte()

	// readFrom演示
	readFrom()

	// newReadWriter演示
	newReadWriter()

	// newScanner演示
	newScanner()

	println()
	println("========================OS Library========================")
}
