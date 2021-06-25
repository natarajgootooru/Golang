package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	// 1) Create a multi-line raw string and print it onto the console.
	var rawStr string = `Hello this is first line
							this is second line
							this is third line`
	fmt.Println(rawStr)

	// 2) create a string & raw string using special characters in it and print it on to the console. Observe the behaviour.
	a := "my !\" do \\ no"
	b := `my !" do \ no`
	fmt.Println(a)
	fmt.Println(b)
	// 3) Find length of a given string and print it on to the console
	fmt.Println("Length of apple: ", len("apple"))
	// 4) Find length of a string which contains few of non-english characters using len() and RuneCountInString() and print it on to the console.
	s := "významu"
	fmt.Printf("významu length: %d\n", len(s))
	fmt.Printf("významu length: %d\n", utf8.RuneCountInString(s))
	// 5) Convert all lowercase letter string into all uppercase.
	fmt.Println(strings.ToUpper("hi john"))
	// 6) Convert all uppercase letter string into all lowercase.
	fmt.Println(strings.ToLower("WELCOME"))
	// 7) Find a substring in a given string and print the result onto the console
	fmt.Println("Contains or not: ", strings.Contains("Welcome to my course", "Course"))
	fmt.Println("Contains or not: ", strings.Contains("Welcome to my course", "course"))
	// 8) Repeat a string 10 times and print it onto the console.
	fmt.Println(strings.Repeat("Hello", 10))
	// 9) Concat two strings.
	fmt.Println("My" + " " + "Computer")
	// 10) Concat a string with an integer
	fmt.Println("My Roll no: " + strconv.Itoa(10))

}
