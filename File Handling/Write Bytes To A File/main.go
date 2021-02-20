package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile(
		"first.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b := []byte("Hello How r u?")
	c, err := f.Write(b)
	fmt.Println("No of bytes written to the file: ", c)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(
		"second.txt",
		[]byte("It is my second example"),
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

}

