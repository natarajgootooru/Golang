package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type contact struct {
	Email  string
	Mobile int
}

type user struct {
	ID   int     `json:"id"`
	Name string  `json:"userName"`
	Job  string  `json:"occupation,omitempty"`
	Cnt  contact `json:"contact"`
}

func main() {

	jsonTxt, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonTxt.Close()
	bytes, _ := ioutil.ReadAll(jsonTxt)
	fmt.Println(string(bytes))

	var users []user
	json.Unmarshal(bytes, &users)

	for _, u := range users {
		fmt.Println("Name: ", u.Name)
		fmt.Println("Profession: ", u.Job)
		fmt.Println(strings.Repeat("*", 25))
	}
}

