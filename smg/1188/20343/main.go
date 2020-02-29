package main

import (
	"fmt"
	"os"
)

func main() {
	var s int
	fmt.Scan(&s)
	if s < 6 || s > 50 {
		fmt.Print("Not in interval")
		os.Exit(1)
	}
	count := 0
	for i := 1; i < s; i++ {
		for j := i + 1; j < s; j++ {
			for k := j + 1; k < s; k++ {
				dig := i + j + k
				if dig == s {
					count++
				}
			}
		}
	}
	fmt.Printf("%d \n", count)
}
