package main

import (
	"commoncommand/user"
	"fmt"
)

func main() {

	fmt.Println("hello go")
	hello := user.Hello()
	fmt.Println(hello)

}
