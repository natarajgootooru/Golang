package main

import "fmt"

type price int

func main() {

	var p price = 100
	var q price = 50
	fmt.Println(p + q)

	var a int = int(p)
	_ = a

	var b price = price(a)
	_ = b
}

