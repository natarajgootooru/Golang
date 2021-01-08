package main

import "fmt"

func main() {

	// Circle Circumference = 2πr
	// π = 3.14
	const pi float64 = 3.14

	var r float64 = 20
	var c float64 = 2 * pi * r
	fmt.Println("Radius:", r, "Circumference:", c)

	r = 30
	c = 2 * pi * r
	fmt.Println("Radius:", r, "Circumference:", c)

	// a := 20
	// const b = 0
	// fmt.Println(a / b)

	const a, b int = 20, 30
	const (
		d float64 = 56
		// e int     = 45
		e
	)
	fmt.Println(d, e)
}
