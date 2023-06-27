package goTime

import (
	"fmt"
	"time"
)

func GoTime() {
	println("========================OS Library========================")
	println()

	now := time.Now()
	fmt.Printf("now: %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T,%T,%T,%T,%T,%T \n", year, month, day, hour, minute, second)

	// 时间戳（毫秒）
	timeStamp := now.Unix()
	fmt.Printf("timeStamp: %v\n", timeStamp)
	// 纳秒
	nano := now.UnixNano()
	fmt.Printf("nano: %v\n", nano)

	// Equal 判断两个时间是否相等
	// Before 判断一个时间是否在另一个时间之前
	// After 判断一个时间是否在另一个时间之后

	// 定时器
	// ticker := time.Tick(time.Second)
	// // ticker是一个通道，可以 range 循环。
	// for t := range ticker {
	// 	fmt.Printf("t: %v\n", t)
	// }

	// 时间格式化
	// 注意：Golang中的格式化时间模版不是 y-m-d H:M:S 而是使用Go的诞生时间2006年1月2号15点04分 为模版的（即 2006 1 2 3 4）
	fmt.Printf("now.Format: %v\n", now.Format("2006-01-02 15:04"))

	// 字符串时间解析为时间对象
	// 加载时区
	location, err := time.LoadLocation("Asia/ShangHai")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		// 按照指定时区和指定格式解析字符串时间
		timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-06-27 22:22:00", location)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("timeObj: %v\n", timeObj)
			fmt.Printf("timeObj.Sub(now): %v\n", timeObj.Sub(now))
		}
	}

	println()
	println("========================OS Library========================")
}
