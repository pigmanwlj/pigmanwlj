package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	for key, num := range nums {
		fmt.Println(key, num)
	}
	kvs := map[string]string{"a": "www", "b": "golang.ltd"}
	for key, num := range kvs {
		fmt.Println(key, num)
	}
}
