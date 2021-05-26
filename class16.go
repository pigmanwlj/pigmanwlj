package main

import (
	"fmt"
)

var GolangCommunity map[int]string

func main() {
	GolangCommunity = make(map[int]string)
	GolangCommunity[1] = "www.Golang.Ltd"
	fmt.Println(GolangCommunity)
}
