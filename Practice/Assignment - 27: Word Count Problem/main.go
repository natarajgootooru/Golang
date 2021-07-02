package main

import (
	"fmt"
	"strings"
)

/*
	Write a program to identify word count in a given string.
*/
func main() {

	input := "we are learning golang i am practicing all examples at my workstation i like the course all examples are so simple"
	parts := strings.Split(input, " ")
	m := make(map[string]int)
	for _, v := range parts {
		if c, ok := m[v]; ok == false {
			m[v] = 1
		} else {
			m[v] = c + 1
		}
	}
	fmt.Printf("%#v\n", m)
}
