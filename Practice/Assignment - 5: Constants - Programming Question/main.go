package main

import "fmt"

/**
Calculate total seconds for the given no of days.
Declare hours per day, minutes per hour and seconds per minute as constants.
Print calculated seconds onto the console.
**/

func main() {

	const HOURS int = 24
	const MINUTES int = 60
	const SECONDS int = 60

	var noOfDays int = 5
	var calculatedSeconds = noOfDays * HOURS * MINUTES * SECONDS
	fmt.Printf("Total Seconds in %d days: %d\n", noOfDays, calculatedSeconds)

	noOfDays = 8
	calculatedSeconds = noOfDays * HOURS * MINUTES * SECONDS
	fmt.Printf("Total Seconds in %d days: %d\n", noOfDays, calculatedSeconds)

	noOfDays = 15
	calculatedSeconds = noOfDays * HOURS * MINUTES * SECONDS
	fmt.Printf("Total Seconds in %d days: %d\n", noOfDays, calculatedSeconds)
}
