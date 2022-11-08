package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lordsalmon/monitormanagement/blacklist"
	"github.com/lordsalmon/monitormanagement/database"
	"github.com/lordsalmon/monitormanagement/env"
	"github.com/lordsalmon/monitormanagement/shell"
	"gopkg.in/robfig/cron.v2"

	log "github.com/sirupsen/logrus"
)

var Blacklist = []string{}

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	log.Info("Monitor Management V3. Made with <3 by Simon Schwedes in Go")
	env.LoadEnv()
	initCronjob()
	fmt.Scanln()
}

func initCronjob() {
	log.Info("Initializing cronjob...")
	cronjob := cron.New()
	log.Info("Interval: ", os.Getenv("INTERVAL"))
	cronjob.AddFunc("@every "+os.Getenv("INTERVAL"), func() {
		blacklist.LoadBlacklist()
		database.InitDatabase()
		windows := shell.GetAllWindows()
		for _, window := range windows {
			ctx := context.Background()
			database.Db.NewInsert().Model(&window).Exec(ctx)
		}
	})
	cronjob.Start()
	log.Info("Cronjob initialized successfully!")
}
