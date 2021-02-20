package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f, err := os.Open("first.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// b := make([]byte, 4)
	// for {
	// 	n, err := f.Read(b)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if n > 0 {
	// 		fmt.Printf("%s", b[:n])
	// 	}
	// }
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}

