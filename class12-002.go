package main

import (
	"fmt"
)

func main() {

	var a int = 10
	var GolangCommunity *int
	fmt.Println("Variable Address: %x", &a)
	GolangCommunity = &a
	fmt.Println("Variable Address: %x", GolangCommunity)
	fmt.Println("Variable Address: %d", *GolangCommunity)

}
