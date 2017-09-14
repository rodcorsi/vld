package vld

import "errors"

type numberVld struct {
	*Validate
	name  string
	value float64
}

func (c *numberVld) Required() *numberVld {
	if c.value == 0 {
		c.err = errors.New("value is required")
	}
	return c
}

func (c *numberVld) Max(max float64) *numberVld {
	if c.err == nil && c.value > max {
		c.err = errors.New("value greater than max")
	}
	return c
}

func (c *numberVld) Min(min float64) *numberVld {
	if c.err == nil && c.value < min {
		c.err = errors.New("value less than min")
	}
	return c
}

func (c *numberVld) Range(min, max float64) *numberVld {
	if c.err == nil && (c.value < min || c.value > max) {
		c.err = errors.New("value must be between min and max")
	}
	return c
}

func (c *numberVld) Ok() bool {
	return c.err == nil
}
