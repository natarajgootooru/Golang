package main

import "fmt"

func main() {

	// 1) Create a slice using make function. Pass slice length as input parameter. Print and observe default zero values.
	var a = make([]int, 3)
	fmt.Printf("%#v | Length: %d | Capacity: %d\n", a, len(a), cap(a))
	// 2) Create a slice using make function by passing desired capacity as input. Print and observe capacity.
	var b = make([]int, 3, 10)
	fmt.Printf("%#v | Length: %d | Capacity: %d\n", b, len(b), cap(b))
	// 3) Copy another slice into above slice and observe if any change in the capacity and when this will change.
	a[0] = 10
	a[1] = 20
	a[2] = 30
	b = append(b, a...)
	fmt.Printf("%#v | Length: %d | Capacity: %d\n", b, len(b), cap(b))
	b = append(b, a...)
	fmt.Printf("%#v | Length: %d | Capacity: %d\n", b, len(b), cap(b))
	b = append(b, a...)
	fmt.Printf("%#v | Length: %d | Capacity: %d\n", b, len(b), cap(b))
	// 4) Create a slice using make function without filled with default zero values.
	var c = make([]int, 0, 10)
	fmt.Printf("%#v | Length: %d | Capacity: %d\n", c, len(c), cap(c))
}
