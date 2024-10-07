package errors2

import (
	"log"
)

func f1(id string) error {
	_, err := getUser(id)
	if err != nil {
		log.Printf("Could not get user %q: %v", id, err)
		return err
	}
	// ...

	return nil
}

func getUser(_ string) (int, error) {
	return 0, nil
}
