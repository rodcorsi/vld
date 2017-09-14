package vld

type Validate struct {
	err error
}

func New() *Validate {
	return &Validate{}
}

func (v *Validate) Err() error {
	return v.err
}

func (v *Validate) String(name, value string) *stringVld {
	return &stringVld{
		Validate: v,
		name:     name,
		value:    value,
	}
}

func (v *Validate) StrPtr(name string, value *string) *stringVld {
	if value == nil {
		return &stringVld{
			Validate: v,
			name:     name,
		}
	}
	return &stringVld{
		Validate: v,
		name:     name,
		value:    *value,
	}
}

func (v *Validate) Number(name string, value float64) *numberVld {
	return &numberVld{
		Validate: v,
		name:     name,
		value:    value,
	}
}

func (v *Validate) NumI64(name string, value int64) *numberVld {
	return &numberVld{
		Validate: v,
		name:     name,
		value:    float64(value),
	}
}
