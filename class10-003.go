package main

import (
	"fmt"
)

var b int = 10

func main() {
	var a, c = 0, 100
	fmt.Println(a, b)
	ret := sum(a, c)
	fmt.Println(ret)
}

func sum(a, b int) int {
	fmt.Println(a, b)
	return a + b
}
