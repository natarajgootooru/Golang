package main

import "fmt"

// 1) Create a new defined type with int as its underlying type
type weight int

func main() {

	// 2) Create a variable to above declared defined type and print its type on to the console.
	var a weight = 20
	fmt.Printf("Variable a type: %T\n", a)
	// 3) Assign above create variable to an integer.
	var b int = int(a)
	fmt.Println(b)
	// 4) Convert an integer type variable to above created defined type.
	var c weight = weight(b)
	_ = c
}
