package main

import (
	"fmt"
)

var golang int

func init() {
	golang = 888
}

func main() {
	if golang == 999 {
	} else {
		fmt.Println("Golang Community:", golang)
	}
}
