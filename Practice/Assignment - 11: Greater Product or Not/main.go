package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Find product of two numbers are greater than given number or not.

This program should take total of 3 input integers from command line. First two numbers will be used to calculate the product. The 3rd input will be used for the comparison. Print the result on to the console.

Example Input: 25, 8, 100
Example Output: the product of 25 & 8 are greater than 100.


Example Input: 50, 10, 1000
Example Output: the product of 50 & 10 are lesser than 1000.

Please make sure you handle all the validations. Error must be handled with in your program.

*/
func main() {

	inputs := os.Args
	if len(inputs) != 4 {
		fmt.Println("Invalid inputs. This program expects total of 3 inputs...")
		return
	}

	a, err := strconv.Atoi(inputs[1])
	if err != nil {
		fmt.Printf("Invalid input: %v. Expected an integer\n", inputs[1])
		return
	}

	b, err := strconv.Atoi(inputs[2])
	if err != nil {
		fmt.Printf("Invalid input: %v. Expected an integer\n", inputs[2])
		return
	}

	c, err := strconv.Atoi(inputs[3])
	if err != nil {
		fmt.Printf("Invalid input: %v. Expected an integer\n", inputs[3])
		return
	}

	if p := a * b; p > c {
		fmt.Printf("The product of %d & %d are greater than %d.\n", a, b, c)
	} else {
		fmt.Printf("The product of %d & %d are lesser than %d.\n", a, b, c)
	}
}
