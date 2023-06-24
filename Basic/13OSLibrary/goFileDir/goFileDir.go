package goFileDir

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	file, err := os.Create("operate/a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("file: %v\n", file)
	}
}

// 创建目录
func createDir() {
	err := os.MkdirAll("operate/a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("目录创建成功")
	}
}

// 删除目录或文件
func removeFileDir() {
	err := os.RemoveAll("operate/a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("目录或文件删除成功")
	}
}

// 获取目录或修改目录
func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}

	os.Chdir(dir + "/operate")
	dir, _ = os.Getwd()
	fmt.Printf("改变后的dir：%v\n", dir)
}

// 重命名文件
func renameFile() {
	err := os.Rename("a.txt", "a1.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("重命名文件成功")
	}
}

// 写文件
func writeFile() {
	err := os.WriteFile("a1.txt", []byte("hello world"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("文件写入成功")
	}
}

// 读取文件
func readFile() {
	file, err := os.ReadFile("a1.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("file: %v\n", string(file[:]))
	}
}

// 打开文件
func openClose() {
	file, _ := os.Open("a1.txt")
	fmt.Printf("file.Name(): %v\n", file.Name())
	file.Close()
}

// 打开文件，如果该文件则创建该文件
func openFileOrCreate() {
	file, _ := os.OpenFile("a2.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	fmt.Printf("file.Name(): %v\n", file.Name())
	file.Close()
}

// 读取文件
func readFileOps() {
	file, _ := os.Open("a2.txt")
	buff := make([]byte, 10)
	// file.Read(buff)
	// 指定读取的偏移量，从第三位开始读取。
	file.ReadAt(buff, 3)
	fmt.Printf("string(buff): %v\n", string(buff))
}

// 读取文件定位（和上面的结果相同）
func readFileSeek() {
	file, _ := os.Open("a2.txt")
	file.Seek(3, 0)
	buff := make([]byte, 10)
	file.Read(buff)
	fmt.Printf("string(buff): %v\n", string(buff))
	file.Close()
}

// 读取目录
func readDirOps() {
	// 首先创建下目录
	os.MkdirAll("c/a", os.ModePerm)
	os.MkdirAll("c/b", os.ModePerm)

	dir, _ := os.ReadDir("c/")
	for _, val := range dir {
		fmt.Printf("val.IsDir(): %v\n", val.IsDir())
		fmt.Printf("val.Name(): %v\n", val.Name())
	}
}

// 写入文件
func writeFileOps() {
	// a3.txt写入一段追加的内容（O_APPEND），如果没有该文件则进行创建（O_CREATE）
	file, err := os.OpenFile("a3.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("a3.txt文件写入成功")
		// file.Write([]byte(" hello golang!!!"))
		// 或者也可：
		file.WriteString("hello Golang~~~")
		file.Close()
	}
}

func writeAtOps() {
	file, err := os.OpenFile("a4.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("a4.txt文件写入成功")
		file.WriteAt([]byte("  leechenze  "), 10)
		file.Close()
	}
}

func GoFileDir() {
	println("========================OS Library========================")
	println()

	// 创建文件
	createFile()

	// 创建目录
	createDir()

	// 删除文件或目录
	removeFileDir()

	// 获取当前目录
	pwd()

	// 重命名文件
	renameFile()

	// 写文件
	writeFile()

	// 读文件
	readFile()

	// 打开文件
	openClose()

	// 打开文件，如果该文件则创建该文件
	openFileOrCreate()

	// 读取文件
	readFileOps()

	// 读取文件定位
	readFileSeek()

	// 读取目录
	readDirOps()

	// 写入文件
	writeFileOps()
	writeAtOps()

	println()
	println("========================OS Library========================")
}
