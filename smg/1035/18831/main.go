package main

import (
	"fmt"
)

func main() {
	var n int
	sum := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var digit int
		fmt.Scan(&digit)
		if digit > 0 {
			sum += digit
		}
	}

	fmt.Println(sum)
}
