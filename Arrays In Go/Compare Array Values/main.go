package main

import "fmt"

func main() {

	a := [3]int{34, 56, 89}
	b := [3]int{34, 56, 890}
	fmt.Println(a == b)
}

