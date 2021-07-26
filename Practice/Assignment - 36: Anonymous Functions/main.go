package main

import "fmt"

func main() {
	// 1) Create a simple anonymous function which declares two integers inside function and prints the product on to the console.
	func() {
		a := 20
		b := 30
		fmt.Println(a * b)
	}()

	// 2) Create a simple anonymous function which takes two integers as inputs and prints the product on to the console.
	func(a int, b int) {
		fmt.Println(a * b)
	}(50, 75)

	// 3) Create a simple anonymous function which takes two integers as inputs and returns the product.
	p := func(a int, b int) int {
		return a * b
	}(50, 75)
	fmt.Println(p)
	// 4) Declare an anonymous function and assign it to a variable. Make a call to the anonymous function using variable reference.
	a := func(a int, b int) int {
		return a * b
	}
	fmt.Println(a(50, 90))
}
