package main

import (
	"fmt"
)

func main() {
	ret := Golang_Ltd(10)
	fmt.Println(ret)
	ret2 := Golang_Unl(20)
	fmt.Println(ret2)
}

func Golang_Ltd(a int) int {
        var c int
        c = a * 2
	return c
}

func Golang_Unl(b int) int {
	return b
}

/*
func Golang_Sum(a, b int) int {
	return a
	return b

	var c int
	c = a * b
	fmt.Println(c)
}
*/
