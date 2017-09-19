package vld

import (
	"regexp"
)

var (
	emailRegex = regexp.MustCompile("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:\\(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22)))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
)

type stringVld struct {
	err   error
	zero  bool
	value string
}

func String(value string) *stringVld {
	return &stringVld{
		value: value,
		zero:  value == "",
	}
}

func StrPtr(value *string) *stringVld {
	if value == nil {
		return &stringVld{
			value: "",
			zero:  true,
		}
	}
	return &stringVld{
		value: *value,
		zero:  false,
	}
}

func (s *stringVld) Required() *stringVld {
	if s.err != nil || !s.zero {
		return s
	}
	s.err = ErrRequired.Gen()
	return s
}

func (s *stringVld) GT(length int) *stringVld {
	if s.err != nil || s.zero || len(s.value) > length {
		return s
	}
	s.err = ErrStringGT.Gen(length)
	return s
}

func (s *stringVld) GTE(length int) *stringVld {
	if s.err != nil || s.zero || len(s.value) >= length {
		return s
	}
	s.err = ErrStringGTE.Gen(length)
	return s
}

func (s *stringVld) LT(length int) *stringVld {
	if s.err != nil || s.zero || len(s.value) < length {
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
	if s.err != nil || s.zero || (len(s.value) >= min && len(s.value) <= max) {
		return s
	}
	s.err = ErrStringLength.Gen(min, max)
	return s
}

func (s *stringVld) Len(length int) *stringVld {
	if s.err != nil || s.zero || len(s.value) == length {
		return s
	}
	s.err = ErrStringLen.Gen(length)
	return s
}

func (s *stringVld) Match(rg *regexp.Regexp) *stringVld {
	if s.err != nil || s.zero || rg.MatchString(s.value) {
		return s
	}
	s.err = ErrStringMatch.Gen()
	return s
}

func (s *stringVld) OneOf(values []string) *stringVld {
	if s.err != nil || s.zero {
		return s
	}
	for _, v := range values {
		if v == s.value {
			return s
		}
	}
	s.err = ErrStringOneOf.Gen(values)
	return s
}

func (s *stringVld) IsEmail() *stringVld {
	if s.err != nil || s.zero || emailRegex.MatchString(s.value) {
		return s
	}
	s.err = ErrStringIsEmail.Gen()
	return s
}

func (s *stringVld) Error() error {
	return s.err
}
