package exceptions

import (
	"fmt"
)

type GRPCError struct {
	StatusCode uint32
	Message    string
}

func (e *GRPCError) Error() string {
	return fmt.Sprintf("status: %d, err: %v", e.StatusCode, e.Message)
}
