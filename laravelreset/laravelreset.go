package laravelreset

import (
	"os/exec"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

// CreateOrReset freshen the database up
func CreateOrReset() {
	artisanExists()
	createDb(getDbName())
	migrate()
}

func artisanExists() {
	_, err := os.Stat("artisan")
	if os.IsNotExist(err) {
		fmt.Println("This is not a Laravel project folder.")
		os.Exit(1)
	}
}

func getDbName() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv("DB_DATABASE")
}

func createDb(db string) {
	createDb := "CREATE DATABASE IF NOT EXISTS " + db + ";"
	create := exec.Command("docker", "exec", "docker_db_1", "mysql", "-u", "homestead", "-psecret", "-e", createDb)
	create.Run()
	fmt.Println("Database '" + db + "' created (if not existed before).")
}

func migrate() {
	migrate := exec.Command("php", "artisan", "migrate:fresh", "--drop-views", "--seed")
	migrate.Stdout = os.Stdout
	migrate.Run()
}