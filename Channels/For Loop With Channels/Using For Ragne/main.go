package main

import "fmt"

func feedChannel(c chan int) {

	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func main() {

	c := make(chan int)
	go feedChannel(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Done")
}

