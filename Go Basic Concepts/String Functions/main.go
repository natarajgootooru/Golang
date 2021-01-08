package main

import (
	"fmt"
	s "strings"
)

func main() {

	var p = fmt.Printf

	// strings.ToUpper converts to upper case.
	p("To upper case: %s\n", s.ToUpper("hello"))

	// strings.ToLower converts to lower case.
	p("To lower case: %s\n", s.ToLower("HELLO"))

	// strings.Contains() verifies substring is part of main string or not
	p("Contains: %t\n", s.Contains("Golang", "Go"))

	// strings.Index() retuns first occurring index of given substring
	p("Index: %d\n", s.Index("Chicken Chicken", "ken"))

	// strings.Repeat() returns a new string consisting of count copies of the input.
	p("Repeat: %s\n", s.Repeat("G", 5))
}
