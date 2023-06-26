package goErrors

import (
	"errors"
	"fmt"
	"time"
)

func GoErrors() {
	println("========================OS Library========================")
	println()

	/** 基础错误 */
	s, err := check("hello")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("s: %v\n", s)
	}

	/** 自定义错误 */
	println()
	err1 := oops()
	if err1 != nil {
		fmt.Printf("err1: %v\n", err1.Error())
		// Error() 方法会自动调用的，和上面的结果相同。
		fmt.Printf("err1: %v\n", err1)
	}

	println()
	println("========================OS Library========================")
}

/** 基础错误 */
func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}

/** 自定义错误 */
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v \n", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}
