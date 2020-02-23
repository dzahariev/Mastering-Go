package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 2 || n > 10 {
		fmt.Println("NO")
		os.Exit(1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}
