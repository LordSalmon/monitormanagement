package blacklist

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var Blacklist []string

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
			Blacklist = append(Blacklist, line)
		}
	}
	fmt.Println("Blacklist loaded successfully!")
	fmt.Println("Blacklist:", Blacklist)
}
