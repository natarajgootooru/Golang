package main

import (
	"fmt"
)

func main() {

	domainMonitors := []domain{
		{name: "https://google.com", email: "alerts@google.com"},
		{name: "https://wikipedia.org", email: "alerts@wikipedia.org"},
		{name: "https://udemy.com", email: "alerts@udemy.com"},
		{name: "https://facebook.com", email: "alerts@facebook.com"},
		{name: "https://youtube.com", email: "alerts@youtube.com"},
	}

	c := make(chan domain)

	for _, d := range domainMonitors {
		go d.getStatus(c)
	}

	for d := range c {
		fmt.Printf("Site: %s | Status:%s\n", d.email, d.status)
		if d.status == "NOT OK" {
			d.sendAlert()
		}
		go d.getStatus(c)
	}
	close(c)
}
