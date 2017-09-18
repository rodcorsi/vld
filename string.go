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

func (s *stringVld) Max(max int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) > max {
		s.err = ErrGreaterThanMax.Gen(max)
	}
	return s
}

func (s *stringVld) Min(min int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) < min {
		s.err = ErrSmallerThanMin.Gen(min)
	}
	return s
}

func (s *stringVld) Length(min, max int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) < min || len(s.value) > max {
		s.err = ErrLength.Gen(min, max)
	}
	return s
}

func (s *stringVld) Len(length int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) != length {
		s.err = ErrLen.Gen(length)
	}
	return s
}

func (s *stringVld) Error() error {
	return s.err
}
