package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Calculate Circle Circumference

Take circle radius input from command line arguments and calculate circle circumference.
*/
func main() {

	const PI float64 = 3.14
	radius, _ := strconv.Atoi(os.Args[1])
	c := 2 * PI * float64(radius)
	fmt.Printf("Circle Circumference: %.2f\n", c)
}
