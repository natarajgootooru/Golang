package main

import "fmt"

func main() {

	// v := sum(10, 20)
	// fmt.Println("Sum: ", v)
	reverseStr("Nataraj")
}

func reverseStr(str string) {
	fmt.Println("Input: ", str)
	for _, v := range []rune(str) {
		defer fmt.Printf("%c", v)
	}
}
func sum(a, b int) (s int) {
	defer done()
	fmt.Println("Inside sum function")
	s = a + b
	return
}

func done() {
	fmt.Println("Inside done function")
}

