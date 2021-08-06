package main

import (
	"fmt"
	"net/http"
	"time"
)

type domain struct {
	name       string
	email      string
	status     string
	statusCode int
}

func (d domain) getStatus(c chan domain) {

	time.Sleep(5 * time.Second)
	resp, err := http.Get(d.name)
	if err != nil {
		fmt.Println(err)
		d.status = "NOT OK"
	} else {
		d.status = "OK"
		d.statusCode = resp.StatusCode
	}
	c <- d
}

func (d domain) sendAlert() {
	fmt.Printf("Alert sent to:%s\n", d.email)
}
