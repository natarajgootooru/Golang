package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	a := "Golang"
	fmt.Printf("Golang length: %d\n", len(a))

	b := "Señor"
	// S takes 1 byte
	// e takes 1 byte,
	// ñ takes 2 bytes
	// o takes 1 byte
	// r takes 1 byte
	fmt.Printf("Señor length: %d\n", len(b))

	fmt.Printf("Señor length: %d\n", utf8.RuneCountInString(b))
}
