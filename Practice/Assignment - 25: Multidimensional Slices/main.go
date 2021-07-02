package main

import (
	"fmt"
	"strings"
)

/*
	Declare a Multidimensional string slices and initialize with string literals.
	Each slice can vary in length.
	1) Traverse through the slice using for-range loop and convert all the elements into uppercase.
	2) Traverse through the slice using simple for loop and convert all the elements into lowercase.
*/
func main() {

	input := [][]string{
		{"ant", "elephant", "cat"},
		{"rose", "jasmine", "Tulip", "Orchid"},
		{"orange", "red", "green", "pink", "black", "white"},
	}
	fmt.Printf("Input: %#v\n", input)
	// 1) Traverse through the slice using for-range loop and convert all the elements into uppercase.
	for i, s := range input {
		for j, v := range s {
			input[i][j] = strings.ToUpper(v)
		}
	}
	fmt.Printf("Uppercase: %#v\n", input)
	// 2) Traverse through the slice using simple for loop and convert all the elements into lowercase.
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			input[i][j] = strings.ToLower(input[i][j])
		}
	}
	fmt.Printf("Lowercase: %#v\n", input)
}
