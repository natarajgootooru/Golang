package main

import "fmt"

func main() {

	s := []int{10, 20, 30}
	t := []int{70, 80, 45, 78, 90}
	n := copy(t, s)
	fmt.Printf("Count: %d | trg: %v\n", n, t)
}

