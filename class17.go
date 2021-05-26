package main

import (
	"fmt"
)

func main() {
	var i uint64 = 15
	ret := GolangLtd(i)
	fmt.Println("%d factorial is %d ", i, ret)
}

func GolangLtd(n uint64) (ret uint64) {

	if n > 0 {
		ret = n * GolangLtd(n-1)
		return ret
	}

	return 1
}
