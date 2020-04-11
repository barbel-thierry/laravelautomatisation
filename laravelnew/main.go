package laravelnew

import "fmt"

// Project initializes a laravel project, if folder does not exist.
func Project(name string) {
	FolderExists(name)

	fmt.Println("Launching Laravel command.")
	Create(name)

	AlterEnv(name)

	fmt.Println("You can now cd into '" + name + "/' before launching 'lpr' to Create database (if not exists).")
}