package laravelopen

import (
	"os"
	"fmt"
	"github.com/pkg/browser"
	"github.com/joho/godotenv"
)

// Run open your current project homepage
// via Laravel Valet and your APP_URL
func Run() {
	browser.OpenURL(getURL())
}

func getURL() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv("APP_URL")
}