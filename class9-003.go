package main

import (
	"fmt"
)

var e int = 50

func main() {
	ret := Golang_Mul(10, 20)
	fmt.Println(ret)
	ret2 := Golang_Sum(10, 20)
	fmt.Println(ret2)
}

func Golang_Mul(a, b int) int {
	return a * b
}

func Golang_Sum(a, b int) int {
	return a + b
}
