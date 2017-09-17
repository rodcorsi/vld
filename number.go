package vld

import "errors"

type numberVld struct {
	Error error
	value float64
}

func Number(value int) *numberVld {
	return &numberVld{
		value: float64(value),
	}
}

func NumF64(value float64) *numberVld {
	return &numberVld{
		value: value,
	}
}

func NumI64(value int64) *numberVld {
	return &numberVld{
		value: float64(value),
	}
}

func (c *numberVld) Required() *numberVld {
	if c.value == 0 {
		c.Error = errors.New("required")
	}
	return c
}

func (c *numberVld) Max(max float64) *numberVld {
	if c.value > max {
		c.Error = errors.New("value greater than max")
	}
	return c
}

func (c *numberVld) Min(min float64) *numberVld {
	if c.value < min {
		c.Error = errors.New("value less than min")
	}
	return c
}

func (c *numberVld) Range(min, max float64) *numberVld {
	if c.value < min || c.value > max {
		c.Error = errors.New("value must be between min and max")
	}
	return c
}
