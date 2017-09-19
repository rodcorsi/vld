package vld

import (
	"fmt"
)

type FieldError interface {
	Field() string
	Message() string
	Error() string
}

type fieldError struct {
	fieldName string
	err       error
}

func (e *fieldError) Error() string {
	return fmt.Sprintf("validation for '%v' failed: %v", e.fieldName, e.err.Error())
}

func (e *fieldError) Field() string {
	return e.fieldName
}

func (e *fieldError) Message() string {
	return e.err.Error()
}
