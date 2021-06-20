package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 1) Convert a float variable to integer and notice precision loss.
	var f float64 = 23.19087
	var e int = int(f)
	fmt.Println(e)
	// 2) Convert an integer into float and print it on to the console.
	var x int = 200
	var y float64 = float64(x)
	fmt.Printf("%f\n", y)
	// 3) Assign an int32 variable to int64 variable and observe the behaviour. Apply type conversion if you notice any compilation errors.
	var a int32 = 20
	var b int64 = int64(a)
	_ = b
	// 4) Try to convert a boolean value into integer and observe the behaviours.
	// boolean value cannot be converted into integer. Type conversion must be applied on related data types only.
	// 5) Observe data truncation behaviours. Assign int64 value to int8 and print them onto the console. Notice data truncation if any.
	var i int64 = 1000
	var j int8 = int8(i)
	fmt.Println(j)
	// 6) Convert an integer to string.
	yearStr := strconv.Itoa(2021)
	fmt.Printf("%q\n", yearStr)
	// 7) Convert string to integer.
	year, _ := strconv.Atoi(yearStr)
	fmt.Printf("%T\n", year)
}
