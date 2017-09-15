package vld

type stringVld struct {
	*Validate
	value string
}

func (v *Validate) String(value string) *stringVld {
	return &stringVld{
		Validate: v,
		value:    value,
	}
}

func (v *Validate) StrPtr(value *string) *stringVld {
	if value == nil {
		return &stringVld{
			Validate: v,
		}
	}
	return &stringVld{
		Validate: v,
		value:    *value,
	}
}

func (c *stringVld) Required() *stringVld {
	if c.value == "" {
		c.AddErrMsg(" value is required")
	}
	return c
}

func (c *stringVld) Max(max int) *stringVld {
	if len(c.value) > max {
		c.AddErrMsg(" value greater than max")
	}
	return c
}

func (c *stringVld) Min(min int) *stringVld {
	if len(c.value) < min {
		c.AddErrMsg(" value less than min")
	}
	return c
}

func (c *stringVld) Length(min, max int) *stringVld {
	if len(c.value) < min || len(c.value) > max {
		c.AddErrMsg(" value must be between min and max")
	}
	return c
}

func (c *stringVld) Len(length int) *stringVld {
	if len(c.value) != length {
		c.AddErrMsg(" value must be between specific length")
	}
	return c
}
