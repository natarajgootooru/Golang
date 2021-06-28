package main

import (
	"fmt"
	"strings"
)

/*
Substring Slice Matching

You will be given two slices as shown below:

a = {"you are my friend", "my region is unknown", "help me in need"}
b = {"are", "region", "need"}

Write a program to compare both slices. You need to validate that slice "b" elements are substring of slice "a" elements at respective indexes. I mean, string a[i] should contain string b[i]. If all elements satisfies this condition print both slices are matching as expected.

Please make sure you will handle all validations like nil checks, length checks, etc.
*/
func main() {
	a := []string{"you are my friend", "my region is unknown", "help me in need"}
	b := []string{"are", "region", "need"}

	isMatching := true
	if len(a) != len(b) {
		isMatching = false
	} else {
		for i, v := range a {
			if strings.Contains(v, b[i]) == false {
				isMatching = false
			}
		}
	}
	if isMatching {
		fmt.Println("Both slices are matching as expected...")
	} else {
		fmt.Println("Both slices are not matching as expected...")
	}
}
