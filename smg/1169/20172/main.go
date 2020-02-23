package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("NO")
		os.Exit(1)
	}

	for i := 1; i <= n; i++ {
		sum := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum != 0 && sum == i {
			fmt.Printf("%d ", i)
		}

	}
}
