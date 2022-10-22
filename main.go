package main

import (
	"fmt"
	"os"

	exec "monitormanagement/exec"

	blacklist "github.com/lordsalmon/monitormanagement/blacklist"
	database "github.com/lordsalmon/monitormanagement/database"

	"github.com/joho/godotenv"
)

var Blacklist = []string{}

func main() {
	fmt.Println("Monitor Management V3. Made with <3 by Simon Schwedes in Go")
	LoadEnv()
	blacklist.LoadBlacklist()
	database.InitDatabase()
	windows := exec.GetAllWindows()
	fmt.Println(windows)
}

func LoadEnv() {
	fmt.Println("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	} else {
		fmt.Println("Environment variables loaded successfully!")
	}
}
