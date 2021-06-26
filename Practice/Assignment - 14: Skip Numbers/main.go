package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Skip Numbers

Write a program to skip numbers while printing numbers in sequence.
You will be printing numbers between 1 to 100.
This program should take an integer input and skip all the numbers that are divisible by the input number.
You must use while loop kind of implementation using for loop.
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

	i := 0
	for i < 100 {
		i++
		if i%num == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}
