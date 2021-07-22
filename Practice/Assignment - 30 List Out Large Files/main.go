package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Insufficient number of inputs found")
	}
	path := os.Args[1]
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%-35s | %s\n", "File Name", "Size")
	fmt.Println("-----------------------------------------------")
	for _, file := range files {
		if file.IsDir() == false && file.Size() > 1024 {
			fmt.Printf("%-35s | %d\n", file.Name(), file.Size())
		}
	}
}
