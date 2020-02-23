package main

import (
	"fmt"
	"sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- calculate(n):
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func calculate(i int) int {
	return i*i + i
}

// https://blog.golang.org/pipelines
func main() {
	done := make(chan struct{})
	defer close(done)

	// Set up the pipeline.
	in := gen(done, 1, 2, 3, 4, 5)

	out := merge(done, in)

	// to consume all the output.
	for i := range out {
		fmt.Println(i)
	}

	// To rea only one value and exit with cancelation
	//fmt.Println(<-out)
}
