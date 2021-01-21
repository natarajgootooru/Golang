package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	i, _ := strconv.Atoi(os.Args[1])
	if i >= 80 {
		fmt.Println("Passed, First Class")
	} else if i >= 60 {
		fmt.Println("Passed, Second Class")
	} else if i >= 40 {
		fmt.Println("Passed, Third Class")
	} else {
		fmt.Println("Failed")
	}
}

