package main

import "fmt"

func main() {

	// 1) Print more than 2 variable separated by comma using println function. Observe how is it printed.
	a, b, c := 10, "Its me", true
	fmt.Println(a, b, c)
	// 2) Print more than 2 variables using print function individually and observe how is it printed.
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
	// 3) Print more than 2 variable using println function individually and observe how is it printed.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// 4) Print an integer using printf function using respective verb with & without line break and observe the behaviour.
	fmt.Printf("Integer value: %d\n", 3000)
	fmt.Printf("Integer value: %d", 4000)
	fmt.Printf("Integer value: %d\n", 5000)
	// 5) Print a float & integer together using printf function.
	fmt.Printf("Integer: %d | Float: %f\n", 4500, 45.25)
	// 6) Try controlling float precision.
	fmt.Printf("Float: %.0f\n", 9465.25)
	fmt.Printf("Float: %.2f\n", 9465.25)
	fmt.Printf("Float: %.4f\n", 9465.25)
	// 7) Print a string & boolean using printf function using respective verb.
	fmt.Printf("String: %s | Boolean: %t\n", "Its golang", true)
	// 8) Try printing type of few variables.
	fmt.Printf("%T | %T | %T\n", a, b, c)
	// 9) Try argument indexing with printf function.
	fmt.Printf("Second Argument: %[2]d | First Argument: %[1]d\n", 30, 60)
}
