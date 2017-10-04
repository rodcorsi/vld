package vld

import "fmt"

const (
	// ErrValidation when validation add some errors
	// Used in String, StrPtr, Number
	ErrValidation = "ErrValidation"

	// ErrRequired when value is required
	ErrRequired = "ErrRequired"

	// ErrStringGT when length is not greater than
	// Used in String, StrPtr
	ErrStringGT = "ErrStringGT"

	// ErrStringGTE when length is not greater or equal than
	ErrStringGTE = "ErrStringGTE"

	// ErrStringLT when length is not smaller than
	ErrStringLT = "ErrStringLT"

	// ErrStringLTE when length is not smaller or equal than
	ErrStringLTE = "ErrStringLTE"

	// ErrStringLength when value is out of length
	ErrStringLength = "ErrStringLength"

	// ErrStringLen when value is out of len
	ErrStringLen = "ErrStringLen"

	// ErrStringMatch when value does not match with regex
	ErrStringMatch = "ErrStringMatch"

	// ErrStringOneOf value is not one of
	ErrStringOneOf = "ErrStringOneOf"

	// ErrStringIsEmail value is not a valid email
	ErrStringIsEmail = "ErrStringIsEmail"

	// ErrNumberRange when value is out of range
	// Used in Number
	ErrNumberRange = "ErrNumberRange"

	// ErrNumberGT when value is not greater than
	ErrNumberGT = "ErrNumberGT"

	// ErrNumberGTE when value is not greater or equal than
	ErrNumberGTE = "ErrNumberGTE"

	// ErrNumberLT when value is not smaller than
	ErrNumberLT = "ErrNumberLT"

	// ErrNumberLTE when value is not smaller or equal than
	ErrNumberLTE = "ErrNumberLTE"
)

var defaultErrMessage = map[string]string{
	ErrValidation:    "validation for '%v' failed: %v",
	ErrRequired:      "required",
	ErrStringGT:      "length is not greater than %v",
	ErrStringGTE:     "length is not greater or equal than %v",
	ErrStringLT:      "length is not smaller than %v",
	ErrStringLTE:     "length is not smaller or equal than %v",
	ErrStringLength:  "out of length(min:%v max:%v)",
	ErrStringLen:     "out of len(%v)",
	ErrStringMatch:   "does not match",
	ErrStringOneOf:   "value is not one of %v",
	ErrStringIsEmail: "value is not a valid email",
	ErrNumberRange:   "out of range(min:%v max:%v)",
	ErrNumberGT:      "value is not greater than %v",
	ErrNumberGTE:     "value is not greater or equal than %v",
	ErrNumberLT:      "value is not smaller than %v",
	ErrNumberLTE:     "value is not smaller or equal than %v",
}

type FormatMessageFunc func(errorID string, args Args) string

var FormatMessage FormatMessageFunc = func(errorID string, args Args) string {
	if format, ok := defaultErrMessage[errorID]; ok {
		return fmt.Sprintf(format, args...)
	}
	return fmt.Sprint(errorID, args)
}
