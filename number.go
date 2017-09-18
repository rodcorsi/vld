package vld

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
	if n.err != nil {
		return n
	}
	if n.value == 0 {
		n.err = ErrRequired.Gen()
	}
	return n
}

func (n *numberVld) Max(max float64) *numberVld {
	if n.err != nil {
		return n
	}
	if n.value > max {
		n.err = ErrGreaterThanMax.Gen(max)
	}
	return n
}

func (n *numberVld) Min(min float64) *numberVld {
	if n.err != nil {
		return n
	}
	if n.value < min {
		n.err = ErrSmallerThanMin.Gen(min)
	}
	return n
}

func (n *numberVld) Range(min, max float64) *numberVld {
	if n.err != nil {
		return n
	}
	if n.value < min || n.value > max {
		n.err = ErrRange.Gen(min, max)
	}
	return n
}

func (n *numberVld) Error() error {
	return n.err
}
