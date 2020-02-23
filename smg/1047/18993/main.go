package main

import (
	"fmt"
)

func main() {
	n := 1
	fmt.Scan(&n)

	if n < 0 || n > 50 {
		fmt.Println("NO")
	}

	fib1 := 1
	//fmt.Printf("%d ", fib1)
	fib2 := 1
	//fmt.Printf("%d ", fib2)
	fib := 2
	//fmt.Printf("%d ", fib)

	for i := 3; i < n; i++ {
		fib1 = fib2
		fib2 = fib
		fib = fib1 + fib2
	}
	fmt.Printf("%d ", fib)
}
