package main

import "fmt"

func main() {
	a := []int{10, 20}
	fmt.Printf("%v - Len: %d - Cap: %d\n", a, len(a), cap(a))
	a = append(a, 30)
	fmt.Printf("%v - Len: %d - Cap: %d\n", a, len(a), cap(a))
	a = append(a, 40)
	fmt.Printf("%v - Len: %d - Cap: %d\n", a, len(a), cap(a))
	a = append(a, 50)
	fmt.Printf("%v - Len: %d - Cap: %d\n", a, len(a), cap(a))
}

