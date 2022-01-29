package clipboard

import (
	"fmt"
	"os/exec"
)

// GetTextFromClipboard - get text from clipboard.
func GetTextFromClipboard() (s string, err error) {
	s, err = Read()

	return s, err
}

func checkCommand(command string) error {
	if _, err := exec.LookPath(command); err != nil {
		return fmt.Errorf("error in checkCommand %w", err)
	}

	return nil
}

// Read reads whatever is in the clipboard, and returns it as a string.
func Read() (string, error) {
	if err := checkCommand("xsel"); err != nil {
		return "Need to install xsel", err
	}

	cmd := exec.Command("xsel")

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error in Read function %w", err)
	}

	return string(out), nil
}
