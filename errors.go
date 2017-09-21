package vld

import "fmt"

const (
	ErrValidation    = "ErrValidation"
	ErrRequired      = "ErrRequired"
	ErrStringGT      = "ErrStringGT"
	ErrStringGTE     = "ErrStringGTE"
	ErrStringLT      = "ErrStringLT"
	ErrStringLTE     = "ErrStringLTE"
	ErrStringLength  = "ErrStringLength"
	ErrStringLen     = "ErrStringLen"
	ErrStringMatch   = "ErrStringMatch"
	ErrStringOneOf   = "ErrStringOneOf"
	ErrStringIsEmail = "ErrStringIsEmail"
	ErrNumberRange   = "ErrNumberRange"
	ErrNumberGT      = "ErrNumberGT"
	ErrNumberGTE     = "ErrNumberGTE"
	ErrNumberLT      = "ErrNumberLT"
	ErrNumberLTE     = "ErrNumberLTE"
)

var defaultErrMessage = map[string]string{
	// Used in String, StrPtr, Number
	// ErrValidation when validation add some errors
	ErrValidation: "validation for '%v' failed: %v",

	// ErrRequired when value is required
	ErrRequired: "required",

	// Used in String, StrPtr
	// ErrStringGT when length is not greater than
	ErrStringGT: "length is not greater than %v",

	// ErrStringGTE when length is not greater or equal than
	ErrStringGTE: "length is not greater or equal than %v",

	// ErrStringLT when length is not smaller than
	ErrStringLT: "length is not smaller than %v",

	// ErrStringLTE when length is not smaller or equal than
	ErrStringLTE: "length is not smaller or equal than %v",

	// ErrStringLength when value is out of length
	ErrStringLength: "out of length(min:%v max:%v)",

	// ErrStringLen when value is out of len
	ErrStringLen: "out of len(%v)",

	// ErrStringMatch when value does not match with regex
	ErrStringMatch: "does not match",

	// ErrStringIn value is not one of
	ErrStringOneOf: "value is not one of %v",

	// ErrStringIsEmail value is not a valid email
	ErrStringIsEmail: "value is not a valid email",

	// Used in Number
	// ErrNumberRange when value is out of range
	ErrNumberRange: "out of range(min:%v max:%v)",

	// ErrNumberGT when value is not greater than
	ErrNumberGT: "value is not greater than %v",

	// ErrNumberGTE when value is not greater or equal than
	ErrNumberGTE: "value is not greater or equal than %v",

	// ErrNumberLT when value is not smaller than
	ErrNumberLT: "value is not smaller than %v",

	// ErrNumberLTE when value is not smaller or equal than
	ErrNumberLTE: "value is not smaller or equal than %v",
}

type FormatMessageFunc func(errorID string, args Args) string

var FormatMessage FormatMessageFunc = func(errorID string, args Args) string {
	if format, ok := defaultErrMessage[errorID]; ok {
		return fmt.Sprintf(format, args...)
	}
	return fmt.Sprint(errorID, args)
}
