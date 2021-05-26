package main

import (
	"fmt"
)

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Println("a+b:", c)
	c = a - b
	fmt.Println("a-b:", c)
	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}
	var d = true
	var f = false

	if d && f {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
