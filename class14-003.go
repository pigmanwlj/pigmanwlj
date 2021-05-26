package main

import (
	"fmt"
)

func main() {
	Golang := []int{1, 2, 3}
	fmt.Println(len(Golang))
	Golang = append(Golang, 4)
	fmt.Println(Golang)
}
