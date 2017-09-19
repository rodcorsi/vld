package vld

type stringVld struct {
	err   error
	value string
}

func String(value string) *stringVld {
	return &stringVld{
		value: value,
	}
}

func (s *stringVld) Required() *stringVld {
	if s.err != nil {
		return s
	}
	if s.value == "" {
		s.err = ErrRequired.Gen()
	}
	return s
}

func (s *stringVld) GT(length int) *stringVld {
	if s.err != nil || len(s.value) > length {
		return s
	}
	s.err = ErrStringGT.Gen(length)
	return s
}

func (s *stringVld) GTE(length int) *stringVld {
	if s.err != nil || len(s.value) >= length {
		return s
	}
	s.err = ErrStringGTE.Gen(length)
	return s
}

func (s *stringVld) LT(length int) *stringVld {
	if s.err != nil || len(s.value) < length {
		return s
	}
	s.err = ErrStringLT.Gen(length)
	return s
}

func (s *stringVld) LTE(length int) *stringVld {
	if s.err != nil || len(s.value) <= length {
		return s
	}
	s.err = ErrStringLTE.Gen(length)
	return s
}

func (s *stringVld) Length(min, max int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) < min || len(s.value) > max {
		s.err = ErrStringLength.Gen(min, max)
	}
	return s
}

func (s *stringVld) Len(length int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) != length {
		s.err = ErrStringLen.Gen(length)
	}
	return s
}

func (s *stringVld) Error() error {
	return s.err
}
