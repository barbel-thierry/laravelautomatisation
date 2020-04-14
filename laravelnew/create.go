package laravelnew

import "os/exec"

// Create launches Laravel 'new' command.
func Create(project string) {
	cmd := exec.Command("laravel", "new", project)
	err := cmd.Run()
	if err != nil {
        panic(err)
    }
}