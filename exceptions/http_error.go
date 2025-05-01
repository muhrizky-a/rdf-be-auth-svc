package exceptions

import (
	"fmt"
)

type HTTPError struct {
	StatusCode uint32
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("status: %d, err: %v", e.StatusCode, e.Message)
}
