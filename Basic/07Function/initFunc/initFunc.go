package main

// 执行顺序：初始化变量 > init > main
var i int = initVar()

func init() {
	println("init is runing")
}

func main() {
	println("main is runing")
}

func initVar() int {
	println("initVar is runing")
	return 0
}
