package main

import "fmt"

func main() {

	arr := [4]int{10, 20, 30, 40}
	a := arr[:]
	fmt.Println(arr, a)
	b := a[0:2]
	c := a[2:]
	fmt.Println(arr, a, b, c)
	a[0] = 50
	fmt.Println(arr, a, b, c)
	// c[0] = 100
	// fmt.Println(a, b, c)
}

