package vld

import "fmt"

type Validate struct {
	errs      Errors
	fieldName string
	message   string
}

func New() *Validate {
	return &Validate{}
}

func NewInit(fieldName string) *Validate {
	return &Validate{
		fieldName: fieldName,
	}
}

func (v *Validate) Init(fieldName string) *Validate {
	v.fieldName = fieldName
	v.message = ""
	return v
}

func (v *Validate) Err() error {
	if v.errs != nil {
		return v.errs
	}
	return nil
}

func (v *Validate) AddErrMsg(message string, a ...interface{}) {
	v.message += fmt.Sprintf(message, a...)
}

func (v *Validate) Ok() bool {
	if v.message != "" {
		v.errs = append(v.errs, &ErrorTag{
			fieldName: v.fieldName,
			message:   v.message,
		})
		return false
	}
	return true
}
