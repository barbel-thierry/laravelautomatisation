package laravelnew

import "os/exec"

// create launches Laravel 'new' command.
func create(project string, version string) {
	cmd := ""
	if version == "" {
		cmd = exec.Command("laravel", "new", project)
	} else {
		cmd = exec.Command("composer", "create-project", "--prefer-dist", "laravel/laravel:" + version,  project)
	}
	err := cmd.Run()
	if err != nil {
        panic(err)
    }
}