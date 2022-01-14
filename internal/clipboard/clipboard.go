package clipboard

import (
	"errors"
	"os/exec"

	"github.com/ZacJoffe/clipboard/xclip"
	"github.com/ZacJoffe/clipboard/xsel"
)

// // Clipboards //
// type Clipboards struct {
// 	Xsel  bool
// 	Xclip bool
// }

// Clipboard //
func Clipboard() (s string, err error) {
	app := checkAppsClipboards()

	if app["xsel"] {
		s, err = xsel.Read()
		if err != nil {
			return "", err
		}
	} else if app["xclip"] {
		s, err = xclip.Read()
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("xsel and xclip not found in system\n")
	}
	return s, err
}

// GetTextFromClipboard //
func GetTextFromClipboard() (s string, err error) {
	s, err = Clipboard()
	return s, err
}

func checkAppsClipboards() (clipboards map[string]bool) {
	clipboards = make(map[string]bool)
	clipboards["xsel"] = false
	clipboards["xclip"] = false

	for k := range clipboards {
		exe := exec.Command("which", k)
		err := exe.Run()
		if err != nil {
			clipboards[k] = false
		} else {
			clipboards[k] = true
			return clipboards
		}
	}
	return clipboards
}
