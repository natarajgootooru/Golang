package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 50}
	b := []int{10, 20, 40}
	r := true

	if len(a) != len(b) {
		r = false
	} else {
		for i, v := range a {
			if v != b[i] {
				r = false
			}
		}
	}

	if r {
		fmt.Println("Both slices are equal")
	} else {
		fmt.Println("Both slices are not equal")
	}
}

