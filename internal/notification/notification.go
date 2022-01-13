package notification

import (
	"github.com/martinlindhe/notify"
)

// Notify //
func Notify(text string) {
	notify.Notify("app name", "Translate", text, "")
}
