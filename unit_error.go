package vld

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
	return FormatMessage(u.errorID, u.args)
}
