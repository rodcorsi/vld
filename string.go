package vld

import (
	"errors"
	"fmt"
)

type stringVld struct {
	err   error
	value string
}

func String(value string) *stringVld {
	return &stringVld{
		value: value,
	}
}

func (c *stringVld) Required() *stringVld {
	if c.value == "" {
		c.err = errors.New("required")
	}
	return c
}

func (c *stringVld) Max(max int) *stringVld {
	if len(c.value) > max {
		c.err = fmt.Errorf("greater than max(%v)", max)
	}
	return c
}

func (c *stringVld) Min(min int) *stringVld {
	if len(c.value) < min {
		c.err = fmt.Errorf("smaller than min(%v)", min)
	}
	return c
}

func (c *stringVld) Length(min, max int) *stringVld {
	if len(c.value) < min || len(c.value) > max {
		c.err = fmt.Errorf("out of length(min:%v max:%v)", min, max)
	}
	return c
}

func (c *stringVld) Len(length int) *stringVld {
	if len(c.value) != length {
		c.err = fmt.Errorf("out of len(%v)", length)
	}
	return c
}

func (c *stringVld) Error() error {
	return c.err
}
