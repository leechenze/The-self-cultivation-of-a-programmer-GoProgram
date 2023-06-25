package goBuiltin

import "fmt"

func GoBuiltin() {
	println("========================OS Library========================")
	println()

	/** 常用函数 */
	// append
	println()
	appendFn()

	// len
	println()
	lenFn()

	// new
	println()
	newFn()

	// make
	println()
	makeFn()

	println()
	println("========================OS Library========================")

}

func makeFn() {
	var i *[]int = new([]int)
	fmt.Printf("i: %v\n", i)
	var i1 []int = make([]int, 10)
	fmt.Printf("i1: %v\n", i1)
}

func newFn() {
	b := new(bool)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %v\n", *b)
	i := new(int)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %v\n", *i)
	s := new(string)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s: %v\n", *s)
}

func lenFn() {
	s := "hello world"
	fmt.Printf("len(s): %v\n", len(s))
	s1 := []int{1, 2, 3}
	fmt.Printf("len(s1): %v\n", len(s1))
}

func appendFn() {
	s := []int{1, 2, 3}
	s1 := append(s, 100)
	fmt.Printf("s1: %v\n", s1)
	s2 := []int{4, 5, 6}
	s3 := append(s1, s2...)
	fmt.Printf("s3: %v\n", s3)
}
