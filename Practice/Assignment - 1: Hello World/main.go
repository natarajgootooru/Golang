package main

import "fmt"

func main() {

	// Print few city names using fmt.Println function
	fmt.Println("Delhi")
	fmt.Println("Bangalore")
	fmt.Println("Mumbai")

	// Print few city names using fmt.Print function and observe the behaviour between Println & Print function.
	fmt.Print("Delhi")
	fmt.Print("Bangalore")
	fmt.Print("Mumbai")

	//Make a call to Print & Println function using multiple values seperated by comma
	fmt.Println()
	fmt.Println("Delhi", "Bangalore", "Mumbai")
	fmt.Print("Delhi", "Bangalore", "Mumbai")

	// Try replacing package declaration postion and observe behaviours.
	// Try removing import statement and observe the behaviour
}
