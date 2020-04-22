package laravelnew

import (
	"os"
	"fmt"
)

// folderExists checks if folder with this Project name already exists.
func folderExists(folder string) {
	_, err := os.Stat(folder)
	
	if !os.IsNotExist(err) {
		fmt.Println("This folder already exists")
		os.Exit(1)
	}
}