package main

import (
	"fmt"
	"strconv"
)

func main() {

	age := 25
	strAge := strconv.Itoa(age)
	_ = strAge
	fmt.Printf("%q\n", strAge)

	strAge = "45"
	age, _ = strconv.Atoi(strAge)
	fmt.Println(age)
}
