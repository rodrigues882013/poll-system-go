package commons

import (
	"fmt"
)

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

type ConflictedError struct {
	Message string
}

func (e ConflictedError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}
