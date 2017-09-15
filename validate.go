package vld

type Validate struct {
	errs      Errors
	fieldName string
	message   string
}

func New() *Validate {
	return &Validate{}
}

func (v *Validate) Init(fieldName string) *Validate {
	v.fieldName = fieldName
	v.message = ""
	return v
}

func (v *Validate) HasError() bool {
	return v.errs != nil
}

func (v *Validate) Err() error {
	return v.errs
}

func (v *Validate) AddErrMsg(message string) {
	v.message += message
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
