package cmdparser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Action is a struct to hold user inputs
type Action struct {
	Command string
	Input   string
}

// ParseInput takes os.Args slice as an input and returns Action
func ParseInput(cmdInput []string) Action {

	var a Action
	// validate does user have provided command line arguments or not.
	if len(cmdInput) == 1 {
		log.Fatal("No inputs are available. The supported actions are \"ListUsers\" & \"UserStats\"")
	}

	switch strings.ToLower(strings.TrimSpace(cmdInput[1])) {
	case "listusers":
		a = Action{Command: "ListUsers"}
	case "userstats":
		if len(cmdInput) != 3 {
			log.Fatal("Invalid input. You must provide user id along with UserStats as an input")
		}
		s := strings.TrimSpace(cmdInput[2])
		if _, err := strconv.ParseInt(s, 10, 64); err != nil {
			fmt.Printf("Invalid User Id: %q.\n", s)
		}
		a = Action{Command: "UserStats", Input: s}
	default:
		log.Fatal("No valid inputs are available. The supported actions are \"ListUsers\" & \"UserStats\"")
	}

	return a
}
