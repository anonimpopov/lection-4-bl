package errors3

import (
	"fmt"
	"os"
)

func f2() error {
	if err, _ := os.Open("settings.txt"); err != nil {
		return fmt.Errorf("could not set up settings: %w", err)
	}

	return nil
}

// could not open settings.txt: open settings.txt: no such file or directory
// could not set up settings: open settings.txt: no such file or directory
