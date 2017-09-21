package vld

import (
	"bytes"
	"errors"
)

type Validate struct {
	errs []FieldError
}

func New() *Validate {
	return &Validate{}
}

func (e *Validate) Ok(fieldName string, err UnitError) bool {
	if err != nil {
		e.errs = append(e.errs, &fieldError{
			fieldName: fieldName,
			err:       err,
		})
		return false
	}
	return true
}

func (e *Validate) Error() error {
	if e.errs == nil {
		return nil
	}
	var buffer bytes.Buffer
	for _, err := range e.errs {
		buffer.WriteString(err.Error())
		buffer.WriteString("\n")
	}
	return errors.New(buffer.String())
}

func (e *Validate) FieldError() []FieldError {
	return e.errs
}
