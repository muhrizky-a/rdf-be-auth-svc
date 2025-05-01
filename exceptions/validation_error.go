package exceptions

import "fmt"

type FailedValidation struct {
	Field      string
	Tag        string
	ParamValue string
}

type ValidationError struct {
	StatusCode  uint32
	Message     string
	Validations []*FailedValidation
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("status: %d, err: %v", e.StatusCode, e.Message)
}
