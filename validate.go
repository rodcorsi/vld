package vld

import (
	"bytes"
)

type Errors []error

func (e Errors) Error() string {
	var buffer bytes.Buffer
	l := len(e)
	if l == 0 {
		return ""
	}
	buffer.WriteString(e[0].Error())
	for i := 1; i < l; i++ {
		buffer.WriteString("\n")
		buffer.WriteString(e[i].Error())
	}
	return buffer.String()
}

type fieldError struct {
	fieldName string
	err       UnitError
}

type Validate struct {
	errs []*fieldError
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
	return e.ErrorFunc(ConstructError)
}

func (e *Validate) ErrorFunc(constructError ContructorErrorFunc) error {
	if e.errs == nil {
		return nil
	}
	var errs Errors
	for _, e := range e.errs {
		uniErr := NewUnitError(ErrValidation, Args{e.fieldName, constructError(e.err)})
		errs = append(errs, constructError(uniErr))
	}
	return errs
}
