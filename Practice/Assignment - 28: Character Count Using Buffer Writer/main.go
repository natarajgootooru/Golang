package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	// create a file using OpenFile function
	f, err := os.OpenFile(
		"practice.txt",
		os.O_WRONLY|os.O_CREATE,
		0644,
	)
	// handle error
	if err != nil {
		log.Fatal(err)
	}
	// close the file once the operations are done
	defer f.Close()
	// create a buffer writer using above created file reference
	bWriter := bufio.NewWriter(f)
	// prepare an input string
	s := "abcdefghijklmnopqrstuvwxyz\n"
	// iterate 1000 times and write the above input into the file
	for i := 1; i <= 1000; i++ {
		// make sure u will add counter to each line
		_, err = bWriter.WriteString(strconv.Itoa(i) + " " + s)
		// handle error
		if err != nil {
			log.Fatal(err)
		}
	}
	// flush the content at the end
	bWriter.Flush()
}
