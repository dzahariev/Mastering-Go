package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Scan(&m)

	var n int
	fmt.Scan(&n)

	r := m % n
	for r != 0 {
		m = n
		n = r
		r = m % r
	}
	fmt.Printf("%d", n)
}
