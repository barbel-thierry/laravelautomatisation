package laravelreset

import (
	"os"
	"fmt"
)

// ArtisanExists checks if artisan file exists, i.e. it is a Laravel project.
func ArtisanExists() {
	_, err := os.Stat("artisan")
	if os.IsNotExist(err) {
		fmt.Println("This is not a Laravel project folder.")
		os.Exit(1)
	}
}