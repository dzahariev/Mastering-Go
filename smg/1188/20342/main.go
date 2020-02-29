package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					dig := i*1000 + j*100 + k*10 + l
					dig2 := i + j*j + k*k*k + l*l*l*l
					if dig == dig2 {
						fmt.Printf("%d \n", dig)
					}
				}
			}
		}
	}
}
