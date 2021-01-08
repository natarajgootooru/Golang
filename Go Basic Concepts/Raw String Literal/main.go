package main

import "fmt"

func main() {

	a := "one two"
	fmt.Println(a)

	b := `one 
	two
	three`
	fmt.Println(b)
	p := "c:\\my dir"
	q := `c:\my dir`
	fmt.Println(p, q)
}
