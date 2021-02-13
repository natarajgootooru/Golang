package main

import "fmt"

func main() {

	s, e := 1280, 2550

	fmt.Printf("\n\n%-7s %-7s %-7s %-10s\n\n",
		"literal", "dec", "hex", "utf encoded")

	for n := s; n <= e; n++ {
		fmt.Printf("%-7c %-7[1]d %-7[1]x % -10x\n", n, string(n))
	}
}

