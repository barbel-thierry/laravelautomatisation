package laravelopen

import (
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

// getURL parses dotenv file to get the defined APP_URL.
func getURL() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv("APP_URL")
}