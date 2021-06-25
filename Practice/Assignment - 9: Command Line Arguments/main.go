package main

import (
	"fmt"
	"os"
)

/*
	Calculate product of two string lengths

	Get two input strings from command line, calculate its length and print its product.
*/
func main() {

	fmt.Println(os.Args)
	a := os.Args[1]
	b := os.Args[2]
	c := len(a) * len(b)
	fmt.Println(c)
}
