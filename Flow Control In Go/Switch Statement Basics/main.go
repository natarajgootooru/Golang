package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi(os.Args[1])

	switch i {
	default:
		fmt.Println("I dont know")
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
}

