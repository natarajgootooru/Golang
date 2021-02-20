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
	delete(m, 4)
	fmt.Printf("%v\n", m)
	// m = make(map[int]string)
	for k := range m {
		delete(m, k)
	}
	fmt.Printf("%#v\n", m)
}

