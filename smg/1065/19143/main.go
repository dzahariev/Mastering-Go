package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1
	max := math.MinInt64
	for n != 0 {
		fmt.Scan(&n)
		if n%2 == 1 && max < n {
			max = n
		}
	}

	if max == math.MinInt64 {
		fmt.Printf("NO")
	} else {
		fmt.Printf("%d", max)
	}
}
