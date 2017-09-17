package vld

import "fmt"

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

func (c *stringVld) Required() *stringVld {
	if c.value == "" {
		c.AddErrMsg(" required")
	}
	return c
}

func (c *stringVld) Max(max int) *stringVld {
	if len(c.value) > max {
		c.AddErrMsg(fmt.Sprintf(" greater than max(%v)", max))
	}
	return c
}

func (c *stringVld) Min(min int) *stringVld {
	if len(c.value) < min {
		c.AddErrMsg(fmt.Sprintf(" smaller than min(%v)", min))
	}
	return c
}

func (c *stringVld) Length(min, max int) *stringVld {
	if len(c.value) < min || len(c.value) > max {
		c.AddErrMsg(fmt.Sprintf(" out of length(min:%v max:%v)", min, max))
	}
	return c
}

func (c *stringVld) Len(length int) *stringVld {
	if len(c.value) != length {
		c.AddErrMsg(fmt.Sprintf(" out of len(%v)", length))
	}
	return c
}
