package main

import "fmt"

func main() {

	a := [3]int{20, 30, 40}
	b := a
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	a[1] = 50
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}

