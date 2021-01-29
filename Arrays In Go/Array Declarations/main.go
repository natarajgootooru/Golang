package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
	var b = [4]int{56, 23, 50, 67}
	fmt.Printf("%#v\n", b)

	c := [3]string{"Raj", "John", "Amir"}
	fmt.Printf("%#v\n", c)

	d := [5]int{45, 60}
	fmt.Printf("%#v\n", d)

	fmt.Printf("index 1: %s\n", c[1])

	var e [5]int
	e[0] = 45
	e[1] = 90
	e[3] = 100
	fmt.Printf("%#v\n", e)

	e[3] = 200
	fmt.Printf("%#v\n", e)

	var f = [...]int{45, 67, 23}
	fmt.Printf("%#v\n", f)

	var g = [3]string{
		"golang",
		"java",
		"python"}

	fmt.Printf("%#v\n", g)
}

