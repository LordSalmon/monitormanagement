package blacklist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var Blacklist []string

func LoadBlacklist() {
	log.Info("Loading blacklist...")
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
	log.Info("Blacklist loaded successfully!")
}
