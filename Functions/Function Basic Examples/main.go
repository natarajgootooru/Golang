package main

import (
	"fmt"
	"strconv"
)

func main() {

	// print()
	// for i := 1; i <= 10; i++ {
	// 	evenOrOdd(i)
	// }
	// m := multiply(10, 20)
	// fmt.Printf("10 * 20 = %d\n", m)
	// m = multiply(15, 20)
	// fmt.Printf("15 * 20 = %d\n", m)
	// i := 1
	// i = increment(i)
	// evenOrOdd(i)
	// i = increment(i)
	// evenOrOdd(i)
	// n, err := toNumber("adfad")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Type: %T\n", n)
	fmt.Printf("Remainder: %d\n", getReminder(24))
	fmt.Printf("Remainder: %d\n", getReminder(89))
}

func getReminder(n int) (r int) {

	r = n % 10
	return // return r
}

func toNumber(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}

func increment(i int) int {
	i++
	return i
}

func multiply(a int, b int) int {

	m := a * b
	return m
}

func evenOrOdd(i int) {
	if i%2 == 0 {
		fmt.Printf("%d is even number\n", i)
	} else {
		fmt.Printf("%d is odd number\n", i)
	}
}
