package main

import "fmt"

func main() {

	const (
		Sunday  = iota + 1 //1
		Monday             //2
		Tuesday            //3
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday,
		Friday, Saturday)

	const (
		Mute = iota
		Mono
		Stereo
		_
		Surround
	)

	fmt.Println(Mute, Mono, Stereo, Surround)
}
