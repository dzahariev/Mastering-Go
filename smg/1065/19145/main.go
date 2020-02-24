package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1
	has := false
	lastdigit := math.MaxInt64
	for n != 0 {
		fmt.Scan(&n)
		if lastdigit < n {
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
