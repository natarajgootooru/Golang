package main

import "fmt"

func main() {

	// 1) Declare a new struct type
	type product struct {
		name     string
		price    float64
		quantity int
	}
	// 2) Declare a variable for the above declared struct type
	p1 := product{name: "Apple", price: 2.25, quantity: 500}
	// 3) Print the above variable onto the console
	fmt.Printf("%#v\n", p1)
	// 4) Declare another variable of the above declared struct type without initializing any of the fields.
	p2 := product{}
	// 5) Print the above variable and notice the default zero values.
	fmt.Printf("%#v\n", p2)
	// 6) Create another variable of same struct type with partial field initializations and observe its field values
	p3 := product{name: "Lifeboy"}
	fmt.Printf("%#v\n", p3)
	// 7) Create another variable of same struct type without fields during initializations and print them onto the console
	p4 := product{"Orange", 2.25, 500}
	fmt.Printf("%#v\n", p4)
	// 8) Try accessing struct values using its fields
	fmt.Printf("Name: %s\n", p1.name)
	fmt.Printf("Name: %f\n", p1.price)
	// 9) Try creating anonymous struct, initialize it and print it onto the console.
	v := struct {
		name    string
		id      int
		address string
	}{
		name:    "Meg",
		id:      1016,
		address: "Bangalore",
	}
	fmt.Printf("%#v\n", v)
}
