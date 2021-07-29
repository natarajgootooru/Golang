package main

import (
	"fmt"
	"strings"
)

type os []string

func (o os) print() {

	for i, v := range o {
		fmt.Printf("Index: %d | Value: %s\n", i, v)
	}
}

func (o os) printInUpperCase() {

	for i, v := range o {
		fmt.Printf("Index: %d | Value: %s\n", i, strings.ToUpper(v))
	}
}
