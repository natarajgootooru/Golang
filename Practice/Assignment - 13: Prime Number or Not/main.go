package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Prime Number or Not

Write a program to identify given number is prime number or not using for loop. You must provide input via command line arguments. Make you do does all input validations. Print the final result on to the console.
*/

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Insufficient number of inputs found")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if num < 2 {
		fmt.Println("Number must be greater than 2...")
		return
	}

	isPrime := true

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Printf("%d is a prime number...\n", num)
	} else {
		fmt.Printf("%d is a non-prime number...\n", num)
	}
}
