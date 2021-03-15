package main

import "fmt"

func main() {

	type Contact struct {
		email  string
		mobile int
	}

	type Employee struct {
		id         int
		name       string
		department string
		contact    Contact
	}

	e1 := Employee{
		id:         1016,
		name:       "Ram",
		department: "Accounts",
		contact: Contact{
			email:  "ram@abc.com",
			mobile: 872202033,
		},
	}

	fmt.Printf("e1: %#v\n", e1)

	fmt.Println("Name: ", e1.name)
	fmt.Println("email: ", e1.contact.email)

	e1.contact.email = "ram_m@abc.com"
	fmt.Println("email: ", e1.contact.email)
}

