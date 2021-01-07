package main

import "fmt"

func main() {
	var a int = 10
	var b = 20
	var c int
	c = b + 20
	b, c = 40, 90

	fmt.Println(a, b, c)
}

