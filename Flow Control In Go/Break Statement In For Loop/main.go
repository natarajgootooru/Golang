package main

import "fmt"

func main() {

	c := 0
	for i := 1; true; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			c++
		}
		if c == 10 {
			break
		}
	}
}

