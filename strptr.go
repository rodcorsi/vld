package vld

import (
	"errors"
	"fmt"
)

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
		s.err = errors.New("required")
	}
	return s
}

func (s *strptrVld) Max(max int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || len(*s.value) > max {
		s.err = fmt.Errorf("greater than max(%v)", max)
	}
	return s
}

func (s *strptrVld) Min(min int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || len(*s.value) < min {
		s.err = fmt.Errorf("smaller than min(%v)", min)
	}
	return s
}

func (s *strptrVld) Length(min, max int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || (len(*s.value) < min || len(*s.value) > max) {
		s.err = fmt.Errorf("out of length(min:%v max:%v)", min, max)
	}
	return s
}

func (s *strptrVld) Len(length int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || (len(*s.value) != length) {
		s.err = fmt.Errorf("out of len(%v)", length)
	}
	return s
}

func (s *strptrVld) Error() error {
	return s.err
}
