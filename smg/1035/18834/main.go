package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
)

func main() {
	var n int
	fmt.Scan(&n)
	min := intsets.MaxInt
	for i := 0; i < n; i++ {
		var digit int
		fmt.Scan(&digit)
		if digit < min {
			min = digit
		}
	}
	fmt.Println(min)
}
