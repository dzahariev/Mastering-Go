package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	prosto := true
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			prosto = false
		}
	}
	if prosto {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
