package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				dig := i*100 + j*10 + k
				dig2 := i + j*j + k*k*k
				if dig == dig2 {
					fmt.Printf("%d \n", dig)
				}
			}
		}
	}
}
