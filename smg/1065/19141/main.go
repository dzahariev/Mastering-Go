package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Scan(&m)

	var n int
	fmt.Scan(&n)

	for m > n || n > m {
		if m > n {
			m = m - n
		} else {
			n = n - m
		}
	}
	fmt.Printf("%d", m)
}
