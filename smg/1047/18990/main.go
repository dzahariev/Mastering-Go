package main

import (
	"fmt"
)

func main() {
	n := 1
	proizwedenie := 1
	for n != 0 {
		fmt.Scan(&n)

		if n != 0 && n%3 == 0 {
			proizwedenie *= n
		}
	}
	if proizwedenie == 1 {
		fmt.Println("NO")
	} else {
		fmt.Printf("%d", proizwedenie)
	}
}
