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

	counter := 0
	suma := 0
	for n != 0 {
		digit := n % 10
		n = n / 10
		counter++
		suma += digit
	}
	fmt.Printf("%d %d", counter, suma)
}
