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

	for n != 0 {
		digit := n % 10
		n = n / 10
		fmt.Print(digit)
	}
}
