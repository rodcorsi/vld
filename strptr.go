package vld

type strptrVld struct {
	err   error
	value *string
}

func StrPtr(value *string) *strptrVld {
	return &strptrVld{
		value: value,
	}
}

func (s *strptrVld) Required() *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil {
		s.err = ErrRequired.Gen()
	}
	return s
}

func (s *strptrVld) Max(max int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || len(*s.value) > max {
		s.err = ErrGreaterThanMax.Gen(max)
	}
	return s
}

func (s *strptrVld) Min(min int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || len(*s.value) < min {
		s.err = ErrSmallerThanMin.Gen(min)
	}
	return s
}

func (s *strptrVld) Length(min, max int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || (len(*s.value) < min || len(*s.value) > max) {
		s.err = ErrLength.Gen(min, max)
	}
	return s
}

func (s *strptrVld) Len(length int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || (len(*s.value) != length) {
		s.err = ErrLen.Gen(length)
	}
	return s
}

func (s *strptrVld) Error() error {
	return s.err
}
