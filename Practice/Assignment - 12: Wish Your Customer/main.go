package main

import (
	"fmt"
	"time"
)

/*
Wish Your Customer

Write a program to wish your customer based on current time using switch statement. You can get current hour of the day using "time.Now().Hour()".

Also identify given day is a weekday or weekend in the same program and print it to the console. You can get current day of the week using "time.Now().Weekday()".
*/

func main() {

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Hi! Good Morning")
	case hour < 17:
		fmt.Println("Hi! Good Afternoon")
	default:
		fmt.Println("Hi! Good Evening")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend buddy...")

	default:
		fmt.Println("It's weekday buddy...")
	}
}
