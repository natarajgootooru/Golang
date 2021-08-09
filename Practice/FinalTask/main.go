package main

import (
	"os"
	"strings"

	"github.com/natarajgootooru/Golang/Practice/FinalTask/cmdparser"
	"github.com/natarajgootooru/Golang/Practice/FinalTask/dao"
)

func main() {

	userAction := cmdparser.ParseInput(os.Args)
	switch strings.ToLower(userAction.Command) {
	case "listusers":
		dao.PrintAllUserOnConsole()
	case "userstats":
		dao.PrintUserStats(userAction.Input)
	}
}
