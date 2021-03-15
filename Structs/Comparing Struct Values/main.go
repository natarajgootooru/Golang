package main

import "fmt"

// Employee struct
type Employee struct {
	name        string
	department  string
	joiningYear int
}

func main() {

	e1 := Employee{name: "John", department: "Accounts", joiningYear: 2015}
	e2 := Employee{name: "John", department: "Accounts", joiningYear: 2015}

	if e1 == e2 {
		fmt.Println("Structs values are equal")
	} else {
		fmt.Println("Structs values are not equal")
	}

	e3 := e1
	fmt.Printf("e3: %#v\n", e3)
	e3.name = "Amar"
	fmt.Printf("e3: %#v\n", e3)
	fmt.Printf("e1: %#v\n", e1)
}

