package goLog

import (
	"fmt"
	"log"
	"os"
)

func GoLog() {
	println("========================OS Library========================")
	println()

	/** print */
	println()
	log.Print("log print")
	log.Printf("log print %d", 100)
	name := "tom"
	age := 20
	log.Print("name + age", " ", name, " ", age)

	/** panic */
	// println()
	// defer println("panic 异常结束后会执行defer")
	// log.Panic("panic log")
	// log.Panicf("panicf log %v", 100)
	// log.Panicln("panicln log")
	// println("end")

	/** fatal */
	// println()
	// defer println("fatal 异常结束后也不会执行defer")
	// log.Fatal("fatal log")
	// log.Fatalf("fatalf log %v", 100)
	// log.Fatalln("fatalln log")

	/** 日志相关配置 */
	println()
	// 配置日志的输出格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 日志的输出前缀
	log.SetPrefix("customPrefix: ")
	// 日志输出到日志文件中
	dir, _ := os.Getwd()
	os.Chdir(dir + "/operate")
	dir, _ = os.Getwd()
	file, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Printf("err: %v", err)
	} else {
		log.SetOutput(file)
		println("日志信息已写入到a.log日志文件中")
	}
	// 自定义logger(将输出格式，输出前缀，输出日志文件组合起来了)
	logger := log.New(file, "customPrefix: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("自定义logger输出")
	// 日志打印
	log.Print("log after setFlags")

	println()
	println("========================OS Library========================")
}
