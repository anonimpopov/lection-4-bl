package errors2

import "fmt"

func f1Good(id string) error {
	_, err := getUser(id)
	if err != nil {
		return fmt.Errorf("get user %q: %w", id, err)
	}
	// ...

	return nil
}
