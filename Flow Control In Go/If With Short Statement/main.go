package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	i, _ := strconv.Atoi(os.Args[1])
	l, _ := strconv.Atoi(os.Args[2])

	var r int

	if r = i * 5; r >= l {
		fmt.Println(r)
		fmt.Println("value is greater than limit")
	} else {
		fmt.Println("value is lesser than limit")
		fmt.Println(r)
	}

	fmt.Println(r)
}

