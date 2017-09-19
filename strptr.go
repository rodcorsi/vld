package vld

import "regexp"

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

func (s *strptrVld) GT(length int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value != nil && len(*s.value) > length {
		return s
	}
	s.err = ErrStringGT.Gen(length)
	return s
}

func (s *strptrVld) GTE(length int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value != nil && len(*s.value) >= length {
		return s
	}
	s.err = ErrStringGTE.Gen(length)
	return s
}

func (s *strptrVld) LT(length int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value != nil && len(*s.value) < length {
		return s
	}
	s.err = ErrStringLT.Gen(length)
	return s
}

func (s *strptrVld) LTE(length int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value != nil && len(*s.value) <= length {
		return s
	}
	s.err = ErrStringLTE.Gen(length)
	return s
}

func (s *strptrVld) Length(min, max int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || (len(*s.value) < min || len(*s.value) > max) {
		s.err = ErrStringLength.Gen(min, max)
	}
	return s
}

func (s *strptrVld) Len(length int) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value == nil || (len(*s.value) != length) {
		s.err = ErrStringLen.Gen(length)
	}
	return s
}

func (s *strptrVld) Match(rg *regexp.Regexp) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value != nil {
		if rg.MatchString(*s.value) {
			return s
		}
	}
	s.err = ErrStringMatch.Gen()
	return s
}

func (s *strptrVld) OneOf(values []string) *strptrVld {
	if s.err != nil {
		return s
	}
	if s.value != nil {
		for _, v := range values {
			if v == *s.value {
				return s
			}
		}
	}
	s.err = ErrStringOneOf.Gen(values)
	return s
}

func (s *strptrVld) Error() error {
	return s.err
}
