package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n < 10000 || n > 99999 {
		fmt.Println("INVALID NUMBER")
		os.Exit(1)
	}

	for i := 1; i <= 5; i++ {
		dig := n % 10
		n = n / 10
		fmt.Print(dig)
	}
}
