package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 40}
	fmt.Printf("%v - Len: %d - Cap: %d\n", a, len(a), cap(a))
	b := a[0:0]
	fmt.Printf("%v - Len: %d - Cap: %d\n", b, len(b), cap(b))
	c := b[0:3]
	fmt.Printf("%v - Len: %d - Cap: %d\n", c, len(c), cap(c))
	d := a[4:4]
	fmt.Printf("%v - Len: %d - Cap: %d\n", d, len(d), cap(d))

}

