package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	monitormanagement "monitormanagement/database"

	"github.com/joho/godotenv"
)

var blacklist = []string{}

func main() {
	fmt.Println("Monitor Management V3. Made with <3 by Simon Schwedes in Go")
	LoadEnv()
	LoadBlacklist()
	monitormanagement.InitDatabase()
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

func LoadBlacklist() {
	fmt.Println("Loading blacklist...")
	requestURL := os.Getenv("BLACKLIST_URL")
	res, err := http.Get(requestURL)
	if err != nil {
		panic(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(resBody), "\n") {
		if len(line) > 0 {
			blacklist = append(blacklist, line)
		}
	}
	fmt.Println("Blacklist loaded successfully!")
	fmt.Println("Blacklist:", blacklist)
}
