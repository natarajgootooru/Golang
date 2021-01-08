package main

import "fmt"

func main() {

	// distance = speed * time
	speed := 100
	time := 5.5
	var distance int

	fmt.Println(speed, time, distance)

	distance = int(float64(speed) * time)
	fmt.Println(distance)

	fmt.Println(int(time))
	fmt.Println(time)

	var temp int32 = int32(speed)
	_ = temp

	// isPossible := true
	// temp = int32(isPossible)
}
