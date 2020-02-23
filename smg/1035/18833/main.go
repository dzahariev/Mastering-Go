package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)
	var n int
	fmt.Scan(&n)
	counter := 0
	for i := 0; i < n; i++ {
		var digit int
		fmt.Scan(&digit)
		if digit == k {
			counter++
		}
	}
	if counter > 0 {
		fmt.Println(counter)
	} else {
		fmt.Println("NO")
	}
}
