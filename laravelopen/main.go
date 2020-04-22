package laravelopen

import "github.com/pkg/browser"

// Run opens your current project homepage
// via Laravel Valet and your APP_URL.
func Run() {
	browser.OpenURL(getURL())
}