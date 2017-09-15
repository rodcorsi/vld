package vld

type numberVld struct {
	*Validate
	value float64
}

func (v *Validate) Number(value int) *numberVld {
	return &numberVld{
		Validate: v,
		value:    float64(value),
	}
}

func (v *Validate) NumF64(value float64) *numberVld {
	return &numberVld{
		Validate: v,
		value:    value,
	}
}

func (v *Validate) NumI64(value int64) *numberVld {
	return &numberVld{
		Validate: v,
		value:    float64(value),
	}
}

func (c *numberVld) Required() *numberVld {
	if c.value == 0 {
		c.AddErrMsg(" value is required")
	}
	return c
}

func (c *numberVld) Max(max float64) *numberVld {
	if c.value > max {
		c.AddErrMsg(" value greater than max")
	}
	return c
}

func (c *numberVld) Min(min float64) *numberVld {
	if c.value < min {
		c.AddErrMsg(" value less than min")
	}
	return c
}

func (c *numberVld) Range(min, max float64) *numberVld {
	if c.value < min || c.value > max {
		c.AddErrMsg(" value must be between min and max")
	}
	return c
}
