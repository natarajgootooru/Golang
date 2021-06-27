package main

import "fmt"

/*
Find Max Sum of Inner Arrays

You will be given a two dimensional integer array.
Iterate over the array one by one and calculate maximum sum of each inner array.
Print only the maximum sum identified during the iteration.
*/
func main() {

	input := [5][3]int{
		{10, 2, 90},
		{78, 90, 1},
		{78, 99, 12},
		{190, 2, 8},
		{66, 99, 88},
	}
	maxSum := 0
	for _, arr := range input {
		sum := 0
		for _, n := range arr {
			sum += n
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	fmt.Printf("Maximum Sum Noticed: %d\n", maxSum)
}
