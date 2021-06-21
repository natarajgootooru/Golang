package main

import "fmt"

func main() {

	// 1) Declare a constant and try reassigning another value. Observe the behaviours.
	// const distance int = 300
	// distance = 300
	// It is not allowed to change constant values once it is declared.

	// 2) Declare uninitialised constant and observe the behaviours.
	// const command string
	// fmt.Println(command)
	// Constants must be initialized durint declarations itself.

	// 3) Try multi constant declarations of same type. Observe is there any compilation error due to unused constants.
	const lower, upper int = 25, 100
	// there will not be any compilation errors due to unused constants. It is not applicable to constants.

	// 4) Try group constant declarations with different types.
	const (
		days    int     = 365
		country string  = "US"
		price   float64 = 34.99
	)

	// 5) Try initializing a constant using a variable & observe the behaviours.
	// var a int = 20
	// const b int = a
	// We cannot initialize constant using a variable. Constants belongs to compile time whereas variable belongs to run time.

	// 6) Try initialising a constant using another constant & observe the behaviours.
	const b int = days * 20
	// we can assing one constant value to another constant which is legal.

	// 7) Declare typeless constants and initialise them using expressions. Print them onto the console.
	const pi = 22 / 7
	fmt.Printf("%T\n", pi)
	fmt.Println(pi)
	const pi_1 = 22 / 7.0
	fmt.Printf("%T\n", pi_1)
	fmt.Println(pi_1)
}
