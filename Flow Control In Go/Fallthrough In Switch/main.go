package main

import "fmt"

func main() {

	i := 45

	switch i {
	case 25:
		fmt.Println(25)
		fallthrough
	case 45:
		fmt.Println(45)
		fallthrough
	case 100:
		fmt.Println(100)

	default:
		fmt.Println(1000)
	}
}

