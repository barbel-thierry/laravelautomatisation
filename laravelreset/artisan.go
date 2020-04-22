package laravelreset

import (
	"os"
	"fmt"
)

// artisanExists checks if artisan file exists, i.e. it is a Laravel project.
func artisanExists() {
	_, err := os.Stat("artisan")
	if os.IsNotExist(err) {
		fmt.Println("This is not a Laravel project folder.")
		os.Exit(1)
	}
}