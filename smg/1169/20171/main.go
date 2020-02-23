package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("NO")
		os.Exit(1)
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			delitel := i
			prost := true
			for j := 2; j <= delitel/2; j++ {
				if delitel%j == 0 {
					prost = false
					break
				}
			}
			if prost {
				fmt.Printf("%d ", delitel)
			}
		}
	}
}
