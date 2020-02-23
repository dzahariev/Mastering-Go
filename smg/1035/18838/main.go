package main

import (
	"fmt"
	"os"
)

func main() {
	var m int
	fmt.Scan(&m)

	if m < 1 || m > 999 {
		fmt.Println("NO")
		os.Exit(1)
	}

	var n int
	fmt.Scan(&n)

	if n < 1 || n > 999 {
		fmt.Println("NO")
		os.Exit(1)
	}

	if m >= n {
		fmt.Println("NO")
		os.Exit(1)
	}

	for i := m; i <= n; i++ {
		digit := i/100*10 + i%10
		if i%digit == 0 {
			fmt.Println(i)
		}
	}
}
