package laravelreset

import (
	"os/exec"
	"os"
)

// migrate runs migration with fresh, seed, and drop views options.
func migrate() {
	migrate := exec.Command("php", "artisan", "migrate:fresh", "--drop-views", "--seed")
	migrate.Stdout = os.Stdout
	err := migrate.Run()
	if err != nil {
		panic(err)
	}
}