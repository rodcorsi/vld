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
	return fmt.Sprintf("validation for '%v' failed%v", e.fieldName, e.message)
}

type Errors []*ErrorTag

func (e Errors) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString(e[0].Error())
	for i := 1; i < len(e); i++ {
		buffer.WriteString("\n")
		buffer.WriteString(e[i].Error())
	}
	return buffer.String()
}
