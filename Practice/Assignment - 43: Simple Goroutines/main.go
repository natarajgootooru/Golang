package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go printNumbers()
	go printAlphabets()
	// time.Sleep(5 * time.Second)
	wg.Wait()
	fmt.Println("Main function executed successfully")
}

func printNumbers() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printAlphabets() {
	defer wg.Done()
	for i := 'a'; i <= 'g'; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
