package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n < 1 || n > 18 {
		fmt.Println("NO")
		os.Exit(1)
	}

	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			if i+j == n {
				fmt.Printf("%d%d", i, j)
				fmt.Println()
			}
		}
	}
}
