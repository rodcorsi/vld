package vld

import "errors"

type numberVld struct {
	err   error
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

func (n *numberVld) Required() *numberVld {
	if n.value == 0 {
		n.err = errors.New("required")
	}
	return n
}

func (n *numberVld) Max(max float64) *numberVld {
	if n.value > max {
		n.err = errors.New("value greater than max")
	}
	return n
}

func (n *numberVld) Min(min float64) *numberVld {
	if n.value < min {
		n.err = errors.New("value less than min")
	}
	return n
}

func (n *numberVld) Range(min, max float64) *numberVld {
	if n.value < min || n.value > max {
		n.err = errors.New("value must be between min and max")
	}
	return n
}

func (n *numberVld) Error() error {
	return n.err
}
