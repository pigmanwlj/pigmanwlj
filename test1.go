package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Printf("Please input password:  \n")
	fmt.Scan(&a)
	if a == 5211314 {
		fmt.Printf("Please input password again: ")
		fmt.Scan(&b)
		if b == 5211314 {
			fmt.Printf("Corret Password, Door will open")
		} else {
			fmt.Printf("Illegal entry, alert")
		}
	} else {
		fmt.Printf("Illegal entry, alert")
	}
}
