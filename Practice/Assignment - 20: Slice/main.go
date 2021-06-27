package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1) Declare an uninitialised slice and print is default zero value.
	var a []int
	fmt.Printf("%#v\n", a)
	// 2) Find length of above uninitialised slice
	fmt.Printf("Length: %d\n", len(a))
	// 3) Try to do index based initialization on uninitialised slice and observe the behaviour.
	// a[0] = 10
	// a[1] = 20
	// It is not possible to do index based initialization when a slice is just declared but not initializes.
	// It throws runtime error.

	// 4) Declare & Initialize a slice. Print it on to the console.
	var b = []string{"Orange", "Red", "Blue", "Green"}
	fmt.Printf("%#v\n", b)
	// 5) Find length of the above slice.
	fmt.Printf("Length: %d\n", len(b))
	// 6) Access & Print elements from the above slice using index.
	fmt.Printf("Index: %d | Element: %s\n", 0, b[0])
	fmt.Printf("Index: %d | Element: %s\n", 1, b[1])
	fmt.Printf("Index: %d | Element: %s\n", 2, b[2])
	// 7) Iterate over the above slice using for loop & range keyword.
	fmt.Println(strings.Repeat("=", 30))
	for i, v := range b {
		fmt.Printf("Index: %d | Element: %s\n", i, v)
	}
}
