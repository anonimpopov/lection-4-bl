package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = &NotFoundError{Name: "test"}

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

func main() {
	err1 := fmt.Errorf("%s: %w", "текст ошибки #1", test())
	err2 := fmt.Errorf("%s: %w", "текст ошибки #2", err1)

	if errors.Is(err2, ErrNotFound) {
		fmt.Println(err2)
	}

	var err *NotFoundError
	if errors.As(err2, &err) {
		fmt.Println(err)
	}
}

func test() error {
	return ErrNotFound
}
