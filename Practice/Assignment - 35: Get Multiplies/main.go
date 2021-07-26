package main

import "fmt"

func main() {

	// variadic function call with one fixed argument
	r1 := getMultiples(8, 2, 9, 6, 7, 12)
	fmt.Println(r1)

	input := []int{10, 15, 20, 26, 30}
	// variadic function call with slice as input
	r2 := getMultiples(2, input...)
	fmt.Println(r2)

}

func getMultiples(factor int, inputs ...int) []int {

	// create a integer slice to store final multiplies
	multiples := make([]int, len(inputs))
	// iterate over slice
	for i, v := range inputs {
		// multiply with factor and store it into slice
		multiples[i] = v * factor
	}
	return multiples
}
