package main

import "fmt"

func main() {
	// 1) Declare a map, print it on to the console and observe the default zero value.
	var a map[int]string
	fmt.Printf("Default Zero Value: %#v\n", a)

	// 2) Initialize the above map and observe the default zero value.
	var b = map[int]string{}
	fmt.Printf("Default Zero Value: %#v\n", b)

	// 3) Add few key-value pairs in to above map.
	b[10] = "Ten"
	b[20] = "Twenty"
	fmt.Printf("Map Values: %#v\n", b)

	// 4) Retrieve values based on index and print it on to the console.
	fmt.Printf("Key: %d | Value: %s\n", 10, b[10])
	fmt.Printf("Key: %d | Value: %s\n", 30, b[30])

	// 5) Update value with existing key, retrieve and print it on to the console.
	b[20] = "TWENTY"
	fmt.Printf("Key: %d | Value: %s\n", 20, b[20])

	// 6) Create a map using make function
	var c = make(map[int]string)
	fmt.Printf("Default Zero Value: %#v\n", c)

	// 7) Declare and initialize a map together
	var d = map[string]int{
		"Ten":    10,
		"Twenty": 20,
		"Thirty": 30,
	}
	fmt.Printf("Map Values: %#v\n", d)

	// 8) Try adding a duplicate key-value pair and observe the behaviour.
	d["Ten"] = 10
	fmt.Printf("Map Values: %#v\n", d)

	// 9) Use for loop and print all key-value pairs.
	for k, v := range d {
		fmt.Printf("Key: %s | Value: %d\n", k, v)
	}

	// 10) Try deleting few pairs.
	delete(d, "Ten")
	fmt.Printf("Map Values: %#v\n", d)

	// 11) Delete all entries from the map.
	for k := range d {
		delete(d, k)
	}
	fmt.Printf("Map Values: %#v\n", d)
}
