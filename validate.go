package vld

type Validate struct {
	errs       Errors
	fieldName  string
	errMessage string
}

func New() *Validate {
	return &Validate{}
}

func (v *Validate) HasError() bool {
	return v.errs != nil
}

func (v *Validate) Err() error {
	return v.errs
}

func (v *Validate) Init(fieldName string) *Validate {
	v.fieldName = fieldName
	v.errMessage = ""
	return v
}

func (v *Validate) Ok() bool {
	if v.errMessage != "" {
		v.errs = append(v.errs, &ErrorTag{
			fieldName: v.fieldName,
			message:   v.errMessage,
		})
		return false
	}
	return true
}
