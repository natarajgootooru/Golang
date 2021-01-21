package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)
	fmt.Printf("Length: %d\n", len(os.Args))
	fmt.Printf("Element at index 0: %s\n", os.Args[0])
	fmt.Printf("Element at index 1: %s\n", os.Args[1])
	fmt.Printf("Element at index 2: %s\n", os.Args[2])
}

