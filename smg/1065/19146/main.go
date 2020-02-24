package main

import (
	"fmt"
)

func main() {
	n := 1
	count := 0
	lastdigit := -1
	maxcount := 0
	for n != 0 {
		fmt.Scan(&n)
		if lastdigit == n {
			count++
		} else {
			if maxcount < count {
				maxcount = count
			}
			count = 0
		}
		lastdigit = n
	}

	fmt.Printf("%d", maxcount)
}
