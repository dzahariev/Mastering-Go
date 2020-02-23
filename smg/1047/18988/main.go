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

	originalNumber := n
	suma := 0
	for n != 0 {
		digit := n % 10
		n = n / 10
		suma += digit
	}

	if originalNumber%suma == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
