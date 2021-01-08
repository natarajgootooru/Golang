package main

import "fmt"

func main() {

	const a = 60 + 3
	const b = 23.3 * a
	const c = "golang" + "-hello"
	const d = 5 > 3
	fmt.Println(a, b, c, d)

	const e = 40
	var f float64 = e
	var g int = e
	var h int32 = e
	var i byte = e
	fmt.Println(f, g, h, i)
}
