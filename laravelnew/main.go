package laravelnew

import "fmt"

// Project initializes a laravel project, if folder does not exist.
func Project(name string) {
	folderExists(name)

	fmt.Println("Launching Laravel command.")
	create(name)

	alterEnv(name)

	fmt.Println("Laravel project has been created.")
}