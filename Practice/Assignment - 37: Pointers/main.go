package main

import "fmt"

func main() {

	// 1) Create a string pointer variable by assigning another string variable.
	var name string = "Golang"
	var namePtr *string = &name
	// 2) Print value, type, address of normal & pointer variable and notice the differences.
	fmt.Printf("name: %-12s | type: %-12T | addr: %-12p\n", name, name, &name)
	fmt.Printf("namePtr: %-12s | type: %-12T | addr: %-12p\n", *namePtr, namePtr, &namePtr)
	// 3) Create a string variable and copy pointer variable value using indirection operator.
	var nameCopy string = *namePtr
	fmt.Printf("nameCopy value: %s\n", nameCopy)
	// 4) Try altering nameCopy value and verify does it affects name value or not.
	nameCopy = "Go"
	fmt.Printf("name: %s | nameCopy: %s\n", name, nameCopy)
	// 5) Try altering namePtr value using indirection operator and observe does it changes name value or not.
	*namePtr = "Its Java Now"
	fmt.Println(name)
	// 6) Create a simple function which takes pointer as input and prints the value. Make a call to that function from main method.
	printName(&name)
	// 7) Create a simple function which returns a pointer back. Make a call to that function from main method.
	s := getName()
	fmt.Println(*s)
}

func printName(i *string) {
	fmt.Println(*i)
}

func getName() *string {

	name := "GOLANG"
	return &name
}
