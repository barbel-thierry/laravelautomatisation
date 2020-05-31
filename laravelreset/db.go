package laravelreset

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

// getDbName parses dotenv file to get the defined DB name.
func getDbName() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv("DB_DATABASE")
}

// createDb logs into docker container to try and create mysql DB.
func createDb(db string, container string) {
	if container == "" {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	createDb := "CREATE DATABASE IF NOT EXISTS `" + db + "`;"
	create := exec.Command("docker", "exec", container, "mysql", "-u", os.Getenv("DB_USERNAME"), "-p"+os.Getenv("DB_PASSWORD"), "-e", createDb)
	create.Run()
	fmt.Println("Database '" + db + "' created (if not existed before).")
}
