package laravelnew

import (
	"os/exec"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

// Project initializes a laravel project, if folder does not exist
func Project(name string) {
	folderExists(name)
	fmt.Println("Launching Laravel command.")
	create(name)
	alterEnv(name)
	fmt.Println("You can now cd into '" + name + "/' before launching 'lpr' to create database (if not exists).")
}

func folderExists(folder string) {
	_, err := os.Stat(folder)
	if !os.IsNotExist(err) {
		fmt.Println("This folder already exists")
		os.Exit(1)
	}
}

func create(project string) {
	cmd := exec.Command("laravel", "new", project)
	cmd.Run()
}

func alterEnv(project string) {
	os.Chdir(project)

	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		fmt.Println(err)
	}

	env["DB_DATABASE"] = project
	env["DB_USERNAME"] = "homestead"
	env["DB_PASSWORD"] = "secret"
	env["APP_URL"] = "http://" + project + ".test"
	env["MAIL_DRIVER"] = "smtp"
	env["MAIL_HOST"] = "localhost"
	env["MAIL_PORT"] = "1025"

	writeErr := godotenv.Write(env, "./.env")
	if writeErr != nil {
		fmt.Println(writeErr)
	}
}