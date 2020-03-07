package main

import "fmt"

func main() {
	x := []int{48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := x[0]
	for index := 1; index < len(x); index++ {
		if smallest > x[index] {
			smallest = x[index]
		}
	}
	fmt.Printf("Smalless %d", smallest)
	fmt.Println()

	fx := make([]float64, len(x))
	for index := 0; index < len(x); index++ {
		fx[index] = float64(x[index])
	}
	fmt.Printf("Average %f", average(fx))
	fmt.Println()

	y := [6]string{"a", "b", "c", "d", "e", "f"}
	z := y[2:5]
	index := 0
	for _, chrz := range z {
		fmt.Printf("z[%d]=%s \n", index, chrz)
		index++
	}

	fmt.Println()

	fmt.Printf("Average2 %f", func() float64 {
		total := 0.0
		for _, value := range fx {
			total += value
		}
		return total / float64(len(fx))
	}())
	fmt.Println()

}

func average(arr []float64) float64 {
	total := 0.0
	for _, value := range arr {
		total += value
	}
	return total / float64(len(arr))
}
