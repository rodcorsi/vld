package vld

type strptrVld struct {
	*Validate
	value *string
}

func (v *Validate) StrPtr(value *string) *strptrVld {
	return &strptrVld{
		Validate: v,
		value:    value,
	}
}

func (c *strptrVld) Required() *strptrVld {
	if c.value == nil {
		c.AddErrMsg(", it is required")
	}
	return c
}

func (c *strptrVld) Max(max int) *strptrVld {
	if c.value == nil || len(*c.value) > max {
		c.AddErrMsg(", greater than max(%v)", max)
	}
	return c
}

func (c *strptrVld) Min(min int) *strptrVld {
	if c.value == nil || len(*c.value) < min {
		c.AddErrMsg(", smaller than min(%v)", min)
	}
	return c
}

func (c *strptrVld) Length(min, max int) *strptrVld {
	if c.value == nil || (len(*c.value) < min || len(*c.value) > max) {
		c.AddErrMsg(", out of length(min:v%, max:v%)", min, max)
	}
	return c
}

func (c *strptrVld) Len(length int) *strptrVld {
	if c.value == nil || (len(*c.value) != length) {
		c.AddErrMsg(", out of len(%v)", length)
	}
	return c
}
