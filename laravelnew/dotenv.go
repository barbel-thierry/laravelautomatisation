package laravelnew

import (
	"github.com/joho/godotenv"
	"os"
)

// AlterEnv edits dotenv file to standardise DB and mail (mailhog) access.
func AlterEnv(project string) {
	os.Chdir(project)

	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		panic(err)
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
		panic(writeErr)
	}
}