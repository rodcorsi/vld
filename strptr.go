package vld

import (
	"errors"
	"fmt"
)

type strptrVld struct {
	Error error
	value *string
}

func StrPtr(value *string) *strptrVld {
	return &strptrVld{
		value: value,
	}
}

func (c *strptrVld) Required() *strptrVld {
	if c.value == nil {
		c.Error = errors.New("required")
	}
	return c
}

func (c *strptrVld) Max(max int) *strptrVld {
	if c.value == nil || len(*c.value) > max {
		c.Error = fmt.Errorf("greater than max(%v)", max)
	}
	return c
}

func (c *strptrVld) Min(min int) *strptrVld {
	if c.value == nil || len(*c.value) < min {
		c.Error = fmt.Errorf("smaller than min(%v)", min)
	}
	return c
}

func (c *strptrVld) Length(min, max int) *strptrVld {
	if c.value == nil || (len(*c.value) < min || len(*c.value) > max) {
		c.Error = fmt.Errorf("out of length(min:%v max:%v)", min, max)
	}
	return c
}

func (c *strptrVld) Len(length int) *strptrVld {
	if c.value == nil || (len(*c.value) != length) {
		c.Error = fmt.Errorf("out of len(%v)", length)
	}
	return c
}
