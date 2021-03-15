package main

import "fmt"

func main() {

	name, department, joiningYear := "John", "Accounts", 2012
	name2, department2, joiningYear2 := "Ram", "Deliver", 2015
	fmt.Println("emp 1: ", name, department, joiningYear)
	fmt.Println("emp 2: ", name2, department2, joiningYear2)

	type employee struct {
		name        string
		department  string
		joiningYear int
	}

	var e1 = employee{
		name:        "John",
		department:  "Accounts",
		joiningYear: 2012,
	}
	fmt.Printf("%#v\n", e1)

	e2 := employee{"Ram", "Delivery", 2015}
	fmt.Printf("%#v\n", e2)

	e1.name = "John M"
	fmt.Println("Name: ", e1.name)
	fmt.Println("Departmebt: ", e1.department)
}

