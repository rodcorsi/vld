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

func (s *stringVld) Required() *stringVld {
	if s.err != nil {
		return s
	}
	if s.value == "" {
		s.err = errors.New("required")
	}
	return s
}

func (s *stringVld) Max(max int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) > max {
		s.err = fmt.Errorf("greater than max(%v)", max)
	}
	return s
}

func (s *stringVld) Min(min int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) < min {
		s.err = fmt.Errorf("smaller than min(%v)", min)
	}
	return s
}

func (s *stringVld) Length(min, max int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) < min || len(s.value) > max {
		s.err = fmt.Errorf("out of length(min:%v max:%v)", min, max)
	}
	return s
}

func (s *stringVld) Len(length int) *stringVld {
	if s.err != nil {
		return s
	}
	if len(s.value) != length {
		s.err = fmt.Errorf("out of len(%v)", length)
	}
	return s
}

func (s *stringVld) Error() error {
	return s.err
}
