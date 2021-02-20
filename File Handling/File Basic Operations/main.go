package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var f *os.File
	var err error
	fmt.Printf("%T\n", f)
	f, err = os.Create("first.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}
	f.Close()

	f1, err := os.OpenFile("first.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	f1.Close()

	var info os.FileInfo
	info, err = os.Stat("first.txt")
	fmt.Println("Name: ", info.Name())
	fmt.Println("Is Directory: ", info.IsDir())
	fmt.Println("Last Modified: ", info.ModTime())
	fmt.Println("File Size: ", info.Size())

	err = os.Remove("first.txt")
	if err != nil {
		log.Fatal(err)
	}
}

