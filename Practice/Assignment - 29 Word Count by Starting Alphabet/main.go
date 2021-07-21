package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	// read input file
	f, err := os.Open("input.txt")
	// handle errors
	if err != nil {
		log.Fatal(err)
	}
	// close the file once the operations are done
	defer f.Close()
	// read all content from the file
	data, err := ioutil.ReadAll(f)
	// handle error
	if err != nil {
		log.Fatal(err)
	}
	// convert byte array to string
	str := string(data[:])
	// convert the string to lower case
	str = strings.ToLower(str)
	// split the string into words
	parts := strings.Split(str, " ")
	// declare a map to hold the result
	r := make(map[string]int)
	// iterate over parts and count the result
	for _, v := range parts {
		c := v[:1]
		r[c] = r[c] + 1
	}
	// print final result
	fmt.Println(r)
}
