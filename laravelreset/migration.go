package laravelreset

import (
	"os/exec"
	"os"
)

// Migrate runs migration with fresh, seed, and drop views options.
func Migrate() {
	migrate := exec.Command("php", "artisan", "migrate:fresh", "--drop-views", "--seed")
	migrate.Stdout = os.Stdout
	migrate.Run()
}