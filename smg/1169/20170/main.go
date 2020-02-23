package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n > 10 {
		fmt.Println("NO")
		os.Exit(1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j+i < n {
				fmt.Printf("%d ", j+1)
			}
		}
		fmt.Println()
	}
}
