package laravelopen

import (
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

// GetURL parses dotenv file to get the defined APP_URL.
func GetURL() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv("APP_URL")
}