package main

import (
	"fmt"
	"strings"
)

func main() {

	a := [5]int{34, 78, 23, 90, 89}
	for i := 0; i < len(a); i++ {
		fmt.Println("Index: ", i, "| Value: ", a[i])
	}
	fmt.Println(strings.Repeat("*", 30))
	for i := range a {
		fmt.Println("Index: ", i, "| Value: ", a[i])
	}
	fmt.Println(strings.Repeat("*", 30))

	for i, v := range a {
		fmt.Println("Index: ", i, "| Value: ", v)
	}
}

