package main

import (
	"fmt"
)

func main() {
	n := 1
	has := false
	lastdigit := 0
	for n != 0 {
		fmt.Scan(&n)
		if lastdigit == n {
			has = true
		}
		lastdigit = n
	}

	if has {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}
}
