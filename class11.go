package main

import (
	"fmt"
)

var GolangCommunity = [10]string{"www.Golang.Ltd", "url"}

func main() {
	for i := 0; i < len(GolangCommunity); i++ {
		fmt.Println(GolangCommunity[i])
	}
}
