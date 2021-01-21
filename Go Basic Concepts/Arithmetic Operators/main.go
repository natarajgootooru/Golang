package main

import "fmt"

func main() {

	a := 5
	b := 3
	r := (a + b) / (a - b) * 2
	// 8 / 2 * 2
	// L -> R -- 4 * 2 = 8
	// R -> L -- 8 / 4 = 2
	fmt.Println(r)

	r = 14 % a
	fmt.Println(r)

	var c float64 = 3.0 / 2
	fmt.Println(c)
}

