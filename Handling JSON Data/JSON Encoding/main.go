package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	users := []user{
		user{1016, "Raj", "Engineer", contact{"raj@abc.com", 12345}},
		user{2014, "John", "Teacher", contact{"john@abc.com", 456786}},
		user{1016, "Abhi", "Police", contact{"abhi@abc.com", 6745890}},
	}
	json, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}

