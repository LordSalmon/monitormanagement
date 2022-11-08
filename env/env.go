package env

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadEnv() {
	log.Info("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", err)
		os.Exit(1)
	} else {
		log.Info("Environment variables loaded successfully!")
	}
}
