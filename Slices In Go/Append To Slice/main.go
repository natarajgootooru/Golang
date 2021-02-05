package main

import "fmt"

func main() {

	a := []int{10, 20, 30}
	a = append(a, 40)
	fmt.Println(a)

	a = append(a, 50, 60, 70)
	fmt.Println(a)

	b := []int{60, 70}
	b = append(b, a...)
	fmt.Println(b)

	var c []int
	c = append(c, 50)
	fmt.Println(c)
}

