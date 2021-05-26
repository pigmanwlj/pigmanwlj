package main

import (
	"fmt"
)

type GolangCommunity struct {
	ID  int
	Url string
}

func main() {
	//      GolangCommunity{ID:100,Url:"www.golang.ltd"}

	fmt.Println(GolangCommunity{ID: 100, Url: "www.golang.ltd"})
	var Golang GolangCommunity
	Golang.ID = 101
	Golang.Url = "www.Golang.Ltd"
	fmt.Println(Golang)
}
