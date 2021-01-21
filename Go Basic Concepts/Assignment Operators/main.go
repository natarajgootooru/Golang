package main

import "fmt"

func main() {

	a, b := 4, 5
	//increment assignment
	a += b // a = a+b
	fmt.Println(a)
	//decrement assignment
	a -= b // a = a - b
	fmt.Println(a)
	//multiplication assignment
	b *= 5 // b = b*5
	fmt.Println(b)
	//division assignment
	b /= 5 // b = b / 5
	fmt.Println(b)
	//modulus assignment
	b %= 3 // b = b % 3
	fmt.Println(b)

	// increment statement
	a++ // a= a+1
	fmt.Println(a)

	// decrement statement
	b-- // b = b-1
	fmt.Println(b)
}

