package vld

type FieldError interface {
	Field() string
	UnitError() UnitError
	Error() string
}

type fieldError struct {
	fieldName string
	err       UnitError
}

func (e *fieldError) Error() string {
	return FormatMessage(ErrValidation, Args{e.fieldName, e.err.Error()})
}

func (e *fieldError) Field() string {
	return e.fieldName
}

func (e *fieldError) UnitError() UnitError {
	return e.err
}
