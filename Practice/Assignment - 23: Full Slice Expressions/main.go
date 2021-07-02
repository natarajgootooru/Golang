package main

import (
	"fmt"
	"strings"
)

/*
	Declare & initialize a slice and create another slice using slice expressions.
	Make sure any addition to the child slice does not affects the source slice.
*/

func main() {

	a := []int{34, 90, 20, 78, 77, 100, 450}
	// creating slice with regular slice expression
	b := a[1:6]
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("b: %#v\n", b)
	// regular slice expression will affect both slices incase of adding new element
	b = append(b, 100000)
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("b: %#v\n", b)
	// creating slice with full slice expression
	fmt.Printf("%s\n", strings.Repeat("=", 50))
	c := a[1:6:6]
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("c: %#v\n", c)
	// full slice expressions avoids element override problems in the source slices due to changes in the child slice
	c = append(c, 897700)
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("c: %#v\n", c)
}
