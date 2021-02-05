package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(a)
	fmt.Println(a[0:3])
	fmt.Println(a[0:4])
	fmt.Println(a[1:4])
	fmt.Println(a[:4])
	fmt.Println(a[:])

	a = a[:len(a)-1]
	fmt.Println(a)
}

