package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

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
