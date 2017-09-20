package vld

type numberVld struct {
	err   error
	value float64
	zero  bool
}

func Number(value int) *numberVld {
	return &numberVld{
		value: float64(value),
		zero:  value == 0,
	}
}

func NumF32(value float32) *numberVld {
	return &numberVld{
		value: float64(value),
		zero:  value == 0,
	}
}

func NumF64(value float64) *numberVld {
	return &numberVld{
		value: value,
		zero:  value == 0,
	}
}

func NumI32(value int32) *numberVld {
	return &numberVld{
		value: float64(value),
		zero:  value == 0,
	}
}

func NumI64(value int64) *numberVld {
	return &numberVld{
		value: float64(value),
		zero:  value == 0,
	}
}

func (n *numberVld) Required() *numberVld {
	if n.err != nil || !n.zero {
		return n
	}
	n.err = ErrRequired.Gen()
	return n
}

func (n *numberVld) GT(value float64) *numberVld {
	if n.err != nil || n.zero || n.value > value {
		return n
	}
	n.err = ErrNumberGT.Gen(value)
	return n
}

func (n *numberVld) GTE(value float64) *numberVld {
	if n.err != nil || n.zero || n.value >= value {
		return n
	}
	n.err = ErrNumberGTE.Gen(value)
	return n
}

func (n *numberVld) LT(value float64) *numberVld {
	if n.err != nil || n.zero || n.value < value {
		return n
	}
	n.err = ErrNumberLT.Gen(value)
	return n
}

func (n *numberVld) LTE(value float64) *numberVld {
	if n.err != nil || n.zero || n.value <= value {
		return n
	}
	n.err = ErrNumberLTE.Gen(value)
	return n
}

func (n *numberVld) Range(min, max float64) *numberVld {
	if n.err != nil || n.zero || (n.value >= min && n.value <= max) {
		return n
	}
	n.err = ErrNumberRange.Gen(min, max)
	return n
}

func (n *numberVld) Error() error {
	return n.err
}
