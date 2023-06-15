package boolean

import "fmt"

func BoolTypeHelloWorld() {
	println("========================Boolean Type start========================")
	println()

	var b1 bool = true
	var b2 bool = false
	var b3 = true
	var b4 = false
	b5 := true
	b6 := false

	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("b4: %v\n", b4)
	fmt.Printf("b5: %v\n", b5)
	fmt.Printf("b6: %v\n", b6)

	println()

	age := 18
	if age >= 20 {
		fmt.Println("Congratulations on reaching adulthood")
	} else {
		fmt.Println("Oh, sorry, you're still under age")
	}

	println()

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i=%v\n", i)
	}

	println()
	println("========================Boolean Type start========================")
}
