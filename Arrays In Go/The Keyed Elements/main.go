package main

import "fmt"

const (
	mon = iota
	tues
	wed
	thur
	fri
	sat
	sun
)

func main() {

	weekend := [7]bool{sat: true, sun: true}
	fmt.Println(weekend)
}

