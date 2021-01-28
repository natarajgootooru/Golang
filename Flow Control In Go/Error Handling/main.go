package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi(os.Args[1]) // 45 23 500 abc
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}

