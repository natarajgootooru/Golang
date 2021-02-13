package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Hello 日本語"
	fmt.Printf("Length In Bytes: %d\n", len(s))
	fmt.Printf("Length In Runes: %d\n", utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("s[%2d] ==> % -10x ==> %q\n", i, string(r), r)
	}

	fmt.Printf("s[%2d] ==> %c\n", 0, s[0])
	fmt.Printf("s[%2d] ==> %c\n", 4, s[4])
	fmt.Printf("s[%2d] ==> %c\n", 6, s[6])
	fmt.Printf("s[%2d] ==> %s\n", 9, s[9:12])

	rns := []rune(s)
	fmt.Printf("Rune Length: %d\n", len(rns))

	for i, r := range rns {
		fmt.Printf("Rune[%d] : %c\n", i, r)
	}
}

