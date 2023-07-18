package goEnviron

import (
	"fmt"
	"os"
)

func GoEnviron() {
	println("========================OS Library========================")
	println()

	// 获取所有的环境变量
	environ := os.Environ()
	for _, val := range environ {
		fmt.Printf("environ.val: %v\n", val)
	}

	// 获取某个环境变量
	gopathEnv := os.Getenv("GOPATH")
	fmt.Printf("gopathEnv: %v\n", gopathEnv)

	// 设置环境变量
	os.Setenv("aaa", "aaa")
	fmt.Printf("os.Getenv(\"aaa\"): %v\n", os.Getenv("aaa"))

	// 查找环境变量
	environ1, b := os.LookupEnv("aaa")
	fmt.Printf("b: %v\n", b)
	fmt.Printf("environ1: %v\n", environ1)

	// 替换环境变量（可以通过${key}获取val）
	os.Setenv("aaa", "aaa")
	os.Setenv("bbb", "bbb")
	fmt.Printf("os.ExpandEnv: %v\n", os.ExpandEnv("${aaa} lives in ${bbb}"))

	println()
	println("========================OS Library========================")
}
