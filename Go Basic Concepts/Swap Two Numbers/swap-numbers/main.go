package main

import "fmt"

func main() {

	var a int = 30
	var b int = 50

	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}

