package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)
	var n int
	fmt.Scan(&n)
	var exists bool
	for i := 0; i < n; i++ {
		var digit int
		fmt.Scan(&digit)
		if digit == k {
			exists = true
		}
	}
	if exists {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
