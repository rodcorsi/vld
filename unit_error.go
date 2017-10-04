package vld

import "fmt"

type Args []interface{}

type UnitError interface {
	ErrorID() string
	Args() Args
	Error() string
}

type unitError struct {
	errorID string
	args    Args
}

func NewUnitError(errorID string, args Args) UnitError {
	return &unitError{
		errorID: errorID,
		args:    args,
	}
}

func (u *unitError) ErrorID() string {
	return u.errorID
}

func (u *unitError) Args() Args {
	return u.args
}

func (u *unitError) Error() string {
	if format, ok := defaultErrMessage[u.errorID]; ok {
		return fmt.Sprintf(format, u.args...)
	}
	return fmt.Sprint(u.errorID, u.args)
}

type ContructorErrorFunc func(unitError UnitError) error

func ConstructErrorDefault(unitError UnitError) error {
	return unitError
}

var ConstructError ContructorErrorFunc = ConstructErrorDefault
