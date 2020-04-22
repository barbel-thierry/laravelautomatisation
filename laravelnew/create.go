package laravelnew

import "os/exec"

// create launches Laravel 'new' command.
func create(project string) {
	cmd := exec.Command("laravel", "new", project)
	err := cmd.Run()
	if err != nil {
        panic(err)
    }
}