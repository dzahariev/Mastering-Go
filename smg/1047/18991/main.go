package main

import (
	"fmt"
)

func main() {
	n := 1
	suma := 0
	fmt.Scan(&n)
	for n > 0 {
		if n%2 == 0 {
			suma += n
		}
		fmt.Scan(&n)
	}
	fmt.Printf("%d", suma)
}
