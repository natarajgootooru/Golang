package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile(
		"first.txt",
		os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bWriter := bufio.NewWriter(f)
	b := []byte("Hello, first input\n")
	bytesWritten, err := bWriter.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("No of bytes Written to the buffer: ", bytesWritten)

	bytesWritten, err = bWriter.WriteString("Hello, this is my second input")
	if err != nil {
		log.Fatal(err)
	}
	bWriter.Flush()

	bWriter.Reset(bWriter)
}

