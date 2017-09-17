package vld

import (
	"errors"
	"fmt"
)

type stringVld struct {
	Error error
	value string
}

func String(value string) *stringVld {
	return &stringVld{
		value: value,
	}
}

func (c *stringVld) Required() *stringVld {
	if c.value == "" {
		c.Error = errors.New("required")
	}
	return c
}

func (c *stringVld) Max(max int) *stringVld {
	if len(c.value) > max {
		c.Error = fmt.Errorf("greater than max(%v)", max)
	}
	return c
}

func (c *stringVld) Min(min int) *stringVld {
	if len(c.value) < min {
		c.Error = fmt.Errorf("smaller than min(%v)", min)
	}
	return c
}

func (c *stringVld) Length(min, max int) *stringVld {
	if len(c.value) < min || len(c.value) > max {
		c.Error = fmt.Errorf("out of length(min:%v max:%v)", min, max)
	}
	return c
}

func (c *stringVld) Len(length int) *stringVld {
	if len(c.value) != length {
		c.Error = fmt.Errorf("out of len(%v)", length)
	}
	return c
}
