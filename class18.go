package main

import (
	"fmt"
)

// define interface
type Igolang interface {
	GET()
	POST()
}

// define struct
type STgolang struct {
}

func (this *STgolang) GET() {

}

func (this *STgolang) POST() {

}

func GolangLTD(i interface{}) {
	_ = i
	fmt.Println(i)
}

func main() {
	// call
	var igolang Igolang
	igolang = new(STgolang)
	_ = igolang
	igolang.GET()
	igolang.POST()
	GolangLTD(102020)
	GolangLTD("www.golang.ltd")
}
