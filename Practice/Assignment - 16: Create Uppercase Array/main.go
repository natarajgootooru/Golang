package main

import (
	"fmt"
	"strings"
)

/*
Create Uppercase Array

Write a program to copy & create an array with all uppercases.
Declare a string array with all lowercase strings. Copy this array to another array by converting all the strings into uppercase.
Finally print the result onto the console.
*/
func main() {

	var inputArr = [...]string{"red", "orange", "blue", "green", "yellow"}
	fmt.Printf("%#v\n", inputArr)
	var copyArr [len(inputArr)]string
	for i, v := range inputArr {
		copyArr[i] = strings.ToUpper(v)
	}
	fmt.Printf("%#v\n", copyArr)
}
