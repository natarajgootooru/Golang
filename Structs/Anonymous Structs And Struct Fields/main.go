package main

import "fmt"

func main() {

	emp := struct {
		name        string
		department  string
		joiningYear int
	}{
		name:        "Ram",
		department:  "Accounts",
		joiningYear: 2015,
	}
	fmt.Printf("%#v\n", emp)

	e1 := struct {
		string
		bool
		int
	}{
		"Ram", false, 2015,
	}
	fmt.Printf("%#v\n", e1)

	e1.string = "john"
	fmt.Printf("%#v\n", e1)

	type employee struct {
		name       string
		department string
		int
	}
	e2 := employee{"Ram", "HR", 2014}
	fmt.Printf("%#v\n", e2)
}

