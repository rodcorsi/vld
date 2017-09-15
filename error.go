package vld

import (
	"bytes"
	"fmt"
)

type ErrorTag struct {
	fieldName string
	message   string
}

func NewErrorTag(fieldName, message string) *ErrorTag {
	return &ErrorTag{
		fieldName: fieldName,
		message:   message,
	}
}

func (e *ErrorTag) Error() string {
	return fmt.Sprintf("%v: %v\n", e.fieldName, e.message)
}

type Errors []*ErrorTag

func (e Errors) Error() string {
	var buffer bytes.Buffer
	for _, et := range e {
		buffer.WriteString(et.Error())
	}
	return buffer.String()
}
