package main

import "fmt"

func main() {

	var src = map[int]string{
		1019: "Shalini",
		1020: "Charan",
	}
	// var src = make(map[int]string)
	src[1016] = "John"
	src[1017] = "Ram"
	src[1018] = "Rahim"
	src[1016] = "Karan"
	fmt.Printf("Map: %#v\n", src)
	fmt.Printf("Length: %d\n", len(src))

	// k := 2016

	// if v, ok := src[k]; ok {
	// 	fmt.Printf("Key: %d | Value: %#v\n", k, v)
	// 	return
	// }
	// fmt.Printf("The key: %d doesnot exists\n", k)

	for k, v := range src {
		fmt.Printf("Key: %d | Value: %#v\n", k, v)
	}
}

