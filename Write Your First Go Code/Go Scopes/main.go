package main

import "fmt"

var s = "hello"

func main() {

	fmt.Println("Its my first go code...", s)
	hello()
	help()
}

func hello() {
	fmt.Println(s)
}
