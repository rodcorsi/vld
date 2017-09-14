package vld

import "errors"

type stringVld struct {
	*Validate
	name  string
	value string
}

func (c *stringVld) Required() *stringVld {
	if c.value == "" {
		c.err = errors.New("value is required")
	}
	return c
}

func (c *stringVld) Max(max int) *stringVld {
	if c.err == nil && len(c.value) > max {
		c.err = errors.New("value greater than max")
	}
	return c
}

func (c *stringVld) Min(min int) *stringVld {
	if c.err == nil && len(c.value) < min {
		c.err = errors.New("value less than min")
	}
	return c
}

func (c *stringVld) Length(min, max int) *stringVld {
	if c.err == nil && (len(c.value) < min || len(c.value) > max) {
		c.err = errors.New("value must be between min and max")
	}
	return c
}

func (c *stringVld) Ok() bool {
	return c.err == nil
}
