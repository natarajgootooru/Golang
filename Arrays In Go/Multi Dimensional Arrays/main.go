package main

import "fmt"

func main() {

	a := [2][3]int{
		{10, 20, 30},
		{15, 25, 35},
	}
	fmt.Println(a)
	sum := 0
	for _, inner := range a {
		for _, v := range inner {
			sum += v
		}
	}
	fmt.Println(sum)
}

