package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		fact := 1
		for j := 1; j <= i; j++ {
			fact *= j
		}
		sum += fact
	}
	fmt.Print(sum)
}
