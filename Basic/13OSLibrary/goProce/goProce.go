package goProce

import (
	"fmt"
	"os"
	"time"
)

func GoProce() {
	println("========================OS Library========================")
	println()

	// 获取当前正在运行的进程ID
	fmt.Printf("os.Getpid()", os.Getpid())
	// 获取父进程ID
	fmt.Printf("os.Getppid()", os.Getppid())
	// 设置新进程的属性
	attr := &os.ProcAttr{
		// files指定新进程继承的活动文件对象
		// 前三个分别为，标准输入，标准输出，标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		// 新进程的环境变量
		Env: os.Environ(),
	}

	// 开启一个新的进程（做个了解吧，此处有错误未能得到解决， TextEdit: no such file or directory）
	process, err := os.StartProcess(
		"TextEdit",
		[]string{
			"/System/Applications/TextEdit.app",
			"/Users/lee/MySkills/The-self-cultivation-of-a-programmer-GoProgram/Basic/13OSLibrary/operate/a4.txt",
		},
		attr,
	)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	println(process)
	fmt.Printf("获取进程ID", process.Pid)

	// 通过进程ID查找进程
	findProcess, _ := os.FindProcess(process.Pid)
	fmt.Printf("findProcess", findProcess)

	// 等待10秒，执行函数
	time.AfterFunc(time.Second*10, func() {
		// 向process进程发送退出信号
		process.Signal(os.Kill)
	})

	// 等待进程process的退出，返回进程状态
	state, err := process.Wait()
	fmt.Printf("state.String(): %v\n", state.String())

	println()
	println("========================OS Library========================")

}
