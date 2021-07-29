package main

import (
	"fmt"
	"strings"
)

func main() {

	osList := os{"Windows", "Linux", "iOS"}
	osList.print()
	fmt.Println(strings.Repeat("*", 40))
	osList.printInUpperCase()
	fmt.Println(strings.Repeat("*", 40))
}
