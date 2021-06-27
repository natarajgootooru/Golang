package main

import "fmt"

/*
Print Array in Reverse

Write a program to print given string array in reverse order.
*/
func main() {

	var inputArr = [5]string{"red", "orange", "blue", "green", "yellow"}
	fmt.Printf("%#v\n", inputArr)
	n := len(inputArr)
	for i := n - 1; i >= 0; i-- {
		fmt.Println(inputArr[i])
	}
}
