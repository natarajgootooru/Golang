package main

import "fmt"

func main() {

	var a [3]int
	fmt.Printf("%#v\n", a)

	b := []int{10, 20, 30}
	fmt.Printf("%#v\n", b)
	fmt.Println(len(b))

	fmt.Println(b == nil)

	fmt.Printf("Index 0: %d\n", b[0])
	fmt.Printf("Index 1: %d\n", b[1])

	for i, v := range b {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}

