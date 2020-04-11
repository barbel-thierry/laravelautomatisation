package laravelreset

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"os/exec"
)

// GetDbName parses dotenv file to get the defined DB name.
func GetDbName() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv("DB_DATABASE")
}

// CreateDb logs into docker container to try and create mysql DB.
func CreateDb(db string, container string) {
	if container == "" {
		container = "docker_db_1"
	}
	createDb := "CREATE DATABASE IF NOT EXISTS " + db + ";"
	create := exec.Command("docker", "exec", container, "mysql", "-u", os.Getenv("DB_USERNAME"), "-p" + os.Getenv("DB_PASSWORD"), "-e", createDb)
	create.Run()
	fmt.Println("Database '" + db + "' created (if not existed before).")
}