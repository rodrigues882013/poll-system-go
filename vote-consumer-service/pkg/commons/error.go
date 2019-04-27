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
