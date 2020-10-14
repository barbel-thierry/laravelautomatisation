package laravelnew

import "fmt"

// Project initializes a laravel project, if folder does not exist.
func Project(name string, version string) {
	folderExists(name)

	fmt.Println("Launching Laravel command.")
	create(name, version)

	alterEnv(name)

	fmt.Println("Laravel project has been created.")
}