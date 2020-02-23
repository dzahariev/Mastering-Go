package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Scan(&m)

	var n int
	fmt.Scan(&n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("X")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}
