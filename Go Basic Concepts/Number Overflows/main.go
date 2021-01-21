package main

import (
	"fmt"
	"math"
)

func main() {

	var a uint8 = 255 //0 to 255
	a++
	fmt.Println(a)

	a--
	fmt.Println(a)

	var b int8 = 127 // -128 to 127
	b++
	fmt.Println(b)

	b--
	fmt.Println(b)

	f := float32(math.MaxFloat32)
	fmt.Println(f * 2)
}

