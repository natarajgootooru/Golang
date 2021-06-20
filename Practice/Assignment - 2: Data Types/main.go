package main

import "fmt"

func main() {
	// 1) Declare variables to all datatypes and print its default zero values on to the console.
	var a int
	var b bool
	var c string
	var d float64

	fmt.Println(a, b, c, d)
	// 2) Initialize above variables with respective data type values and print them on to the console.
	var a1 int = 20
	var b1 bool = true
	var c1 string = "Apple"
	var d1 float64 = 12.45

	fmt.Println(a1, b1, c1, d1)
	// 3) Declare any datatype variable names with special characters (!, #, $, %, etc) in it and observe the behaviors.
	// var a!nt int
	/*
		Special characters are not allowed as part of variable names.
	*/
	// 4) Create a variable name starting with number and observe the behaviour.
	//var 2count int
	// Any variable name must not start with a number.
	// 5) Create a variable name with a number in between and observe the behaviour.
	var count2 int
	_ = count2
	//  a variable name can contains a number in between the characters but it is not allowed to create a variable name starting with number.
	// 6) Print your name using a string literal.
	fmt.Println("Nataraja Gootooru")
	// 7) Print a non-english sentence using a string literal.
	fmt.Println("这是一份非常简单的说明书")
	// 8) Create a variable name with go reserve keywords and observe the behaviour.
	// var continue string
	// reserve keywards are not allowed as variable names
	// 9) Just declare any variable and compile the code without using the variable anywhere. Observe the behaviour. And then print the variable and notice the behaviour once again.
	var country string
	fmt.Println(country)
	// unused characters will throw compilation error either in block or method scope.
	// 10) Declare an unused variable and apply blank identifier. Observe the behaviour.
	var state string
	_ = state
	// you can assign any data type variables to blank identifier.
	// 11) Try multivariable declarations & initializations.
	var (
		name   string = "Ram"
		age    int    = 25
		mobile int    = 3412897845
	)
	fmt.Println(name, age, mobile)
	// 12) Try & observe couple of type inference based initializations.
	var gender = "male"
	var pincode = 560103
	var enabled = false
	fmt.Println(gender, pincode, enabled)
	// 13) Try short declaration for an int variable and print the value onto the console.
	hours := 14
	fmt.Println(hours)
	// 14) Try short declaration for a string variable and print the value onto the console.
	day := "Sunday"
	fmt.Println(day)
	// 15) Try short declarations at package scope and observe the behaviour.
	// short declarations are only allowed at method or block scope not at package scope.
	// 16) Try couple of multi-variable declarations.
	var year, date int = 2021, 21
	fmt.Println(year, date)
	var isWeekday, orderNumber = false, 34213
	fmt.Println(isWeekday, orderNumber)
	// 17) Try reassign values to already declared variables and print them onto the console.
	isWeekday = true
	fmt.Println(isWeekday)
}
