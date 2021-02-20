package main

import "fmt"

func main() {

	m := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
	}
	fmt.Printf("%v\n", m)

	c := make(map[int]string)
	for k, v := range m {
		c[k] = v
	}
	fmt.Printf("Map m: %v\nMap c: %v\n", m, c)

	c[5] = "Five"
	fmt.Printf("Map m: %v\nMap c: %v\n", m, c)
}

