package main

import (
	"fmt"
)

func main() {
	n := 1
	counter := 0
	for n != 0 {
		fmt.Scan(&n)

		if n < 0 {
			counter++
		}
	}

	fmt.Printf("%d", counter)
}
