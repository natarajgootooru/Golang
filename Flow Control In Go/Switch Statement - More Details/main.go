package main

import "fmt"

func main() {

	// switch s := os.Args[1]; s {
	// case "Yellow", "Red", "Green":
	// 	fmt.Println("Color")
	// case "India", "USA":
	// 	fmt.Println("Country")
	// default:
	// 	fmt.Println("I dont know")
	// }

	i := 0
	switch {
	case i > 0:
		fmt.Println("Positive")
	case i < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}

