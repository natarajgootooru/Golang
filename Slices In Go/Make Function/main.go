package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 40}
	var b = make([]int, 0, 2*len(a))
	fmt.Printf("%v | Len: %d | Cap: %d\n", b, len(b), cap(b))
	for _, v := range a {
		b = append(b, v)
		fmt.Printf("%v | Len: %d | Cap: %d\n", b, len(b), cap(b))
	}
}

