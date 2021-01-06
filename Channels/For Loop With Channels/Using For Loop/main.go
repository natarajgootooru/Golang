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

	for {
		v, ok := <-c
		if ok {
			fmt.Println(v, ok)
		} else {
			fmt.Println(v, ok, "Channel is closed")
			break
		}
	}

	fmt.Println("Done")
}

