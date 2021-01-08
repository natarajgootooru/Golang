package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := "Apple"
	b := "Orange"
	c := a + " " + b
	fmt.Println(c)

	c += " Grape" + strconv.Itoa(10)
	fmt.Println(c)
}
