package main

import (
	"fmt"

	"github.com/lordsalmon/monitormanagement/blacklist"
	"github.com/lordsalmon/monitormanagement/database"
	"github.com/lordsalmon/monitormanagement/env"
	"github.com/lordsalmon/monitormanagement/shell"
)

var Blacklist = []string{}

func main() {
	fmt.Println("Monitor Management V3. Made with <3 by Simon Schwedes in Go")
	env.LoadEnv()
	blacklist.LoadBlacklist()
	database.InitDatabase()
	windows := shell.GetAllWindows()
	fmt.Println(windows)
}
