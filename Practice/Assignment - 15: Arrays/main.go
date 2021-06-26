package main

import "fmt"

func main() {

	// 1) Declare an int array and print it on to the console. Observe the default zero values.
	var a [5]int
	fmt.Println(a)
	// 2) Print the above array type & default zero values together.
	fmt.Printf("%#v\n", a)
	// 3) Declare & initialize string array and print it on to the console.
	var s = [3]string{"apple", "orange", "grapes"}
	fmt.Printf("%#v\n", s)
	s1 := [3]string{"green", "red", "blue"}
	fmt.Printf("%#v\n", s1)
	// 4) Declare an int array and try partial initialization. Print & observe the values.
	b := [5]int{10, 20}
	fmt.Printf("%#v\n", b)
	// 5) Try reading array elements using its index and print it on to the console.
	fmt.Printf("Element at index 0: %s\n", s[0])
	fmt.Printf("Element at index 1: %s\n", s[1])
	fmt.Printf("Element at index 2: %s\n", s[2])
	// 6) Declare and initialize array elements using its index
	var c [5]int
	c[0] = 20
	c[3] = 50
	fmt.Printf("%#v\n", c)
	// 7) Try creating an array using ellipses and print it on to the console.
	d := [...]int{90, 25, 89, 1098}
	fmt.Printf("%#v\n", d)
	// 8) Try array multiline initialization.
	var e = [3]float64{
		30.50,
		25.90,
		90.99,
	}
	fmt.Printf("%#v\n", e)
}
