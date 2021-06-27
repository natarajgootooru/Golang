package main

import "fmt"

func main() {

	// 1) Create a float64 array using keyed indexes and print it onto the console.
	a := [4]float64{0: 10.99, 1: 99.99, 2: 56.09, 3: 21.50}
	fmt.Println(a)
	// 2) Change the order of the index in the above array, print it onto the console and observe the element order.
	b := [4]float64{2: 56.09, 3: 21.50, 0: 10.99, 1: 99.99}
	fmt.Println(b)
	// 3) Try keyed indexing using ellipses and observe the printed array on the console
	c := [...]string{0: "Apple", 1: "Jack", 2: "Orange"}
	fmt.Printf("%#v\n", c)
	// 4) Try partial elements using keyed indexing and observe the printed array on the console
	d := [5]string{0: "Apple", 3: "Jack"}
	fmt.Printf("%#v\n", d)
	// 5) Replace length with ellipses in the above array and observe the printed array on the console
	e := [...]string{0: "Apple", 3: "Jack"}
	fmt.Printf("%#v\n", e)
	// 6) Try mix of with keyed & without keyed indexing and observe the printed array on the console
	f := [...]string{0: "Apple", 3: "Jack", "Orange"}
	fmt.Printf("%#v\n", f)
}
