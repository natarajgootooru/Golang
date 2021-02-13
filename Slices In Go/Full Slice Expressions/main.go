package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}
	fmt.Printf("%v - Len: %d - Cap: %d\n", a, len(a), cap(a))

	b := a[0:2:2]
	fmt.Printf("%v - Len: %d - Cap: %d\n", b, len(b), cap(b))

	b = append(b, 50)
	fmt.Println(a, b)
}

