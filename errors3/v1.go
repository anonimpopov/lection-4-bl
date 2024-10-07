package errors3

import (
	"fmt"
	"os"
)

func f1() error {
	if err, _ := os.Open("settings.txt"); err != nil {
		return fmt.Errorf("could not open settings.txt: %w", err)
	}

	return nil
}
