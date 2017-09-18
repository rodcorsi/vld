package vld

import "bytes"

type Validate struct {
	errs []FieldError
}

func New() *Validate {
	return &Validate{}
}

func (e *Validate) Ok(fieldName string, err error) bool {
	if err != nil {
		e.errs = append(e.errs, &fieldError{
			fieldName: fieldName,
			err:       err,
		})
		return false
	}
	return true
}

func (e *Validate) HasError() bool {
	return len(e.errs) > 0
}

func (e *Validate) Error() string {
	var buffer bytes.Buffer
	for _, err := range e.errs {
		buffer.WriteString(err.Error())
		buffer.WriteString("\n")
	}
	return buffer.String()
}
