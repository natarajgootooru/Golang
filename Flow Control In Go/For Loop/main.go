package main

import "fmt"

func main() {

	n := 100
	sum := 0
	i := 1

	for i <= n {
		sum += i // sum = sum + i
		i++
	}

	fmt.Println(sum)
}

