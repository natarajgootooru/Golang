package main

import "fmt"

func main() {

	// 1) Append an element to a slice
	a := []int{90, 56, 89, 7}
	a = append(a, 40)
	fmt.Printf("%#v\n", a)

	// 2) Append multiple elements to a slice
	a = append(a, 80, 4, 908)
	fmt.Printf("%#v\n", a)

	// 3) Append one slice to another slice
	b := []int{999, 888, 777}
	a = append(a, b...)
	fmt.Printf("%#v\n", a)

	// 4) Extract first 4 elements of a slice.
	c := []int{90, 56, 89, 7, 999, 888, 777}
	fmt.Printf("%#v\n", c[0:4])

	// 5) Extract 2 elements from index 3.
	fmt.Printf("%#v\n", c[3:5])

	// 6) Extract last 5 elements.
	fmt.Printf("%#v\n", c[len(c)-5:])

	// 7) Extract all the elements except last two elements.
	fmt.Printf("%#v\n", c[:len(c)-2])
}
