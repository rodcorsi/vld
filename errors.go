package vld

// Used in String, StrPtr, Number
var (
	// ErrRequired when value is required
	ErrRequired ErrorGen = NewErrorGen("required")
)

// used in String, StrPtr
var (
	// ErrStringGT when length is not greater than
	ErrStringGT ErrorGen = NewErrorGen("length is not greater than %v")

	// ErrStringGTE when length is not greater or equal than
	ErrStringGTE ErrorGen = NewErrorGen("length is not greater or equal than %v")

	// ErrStringLT when length is not smaller than
	ErrStringLT ErrorGen = NewErrorGen("length is not smaller than %v")

	// ErrStringLTE when length is not smaller or equal than
	ErrStringLTE ErrorGen = NewErrorGen("length is not smaller or equal than %v")

	// ErrStringLength when value is out of value
	ErrStringLength ErrorGen = NewErrorGen("out of value(min:%v max:%v)")

	// ErrStringLen when value is out of len
	ErrStringLen ErrorGen = NewErrorGen("out of len(%v)")
)

// used in Number
var (
	// ErrNumberRange when value is out of range
	ErrNumberRange ErrorGen = NewErrorGen("out of range(min:%v max:%v)")

	// ErrNumberGT when value is not greater than
	ErrNumberGT ErrorGen = NewErrorGen("value is not greater than %v")

	// ErrNumberGTE when value is not greater or equal than
	ErrNumberGTE ErrorGen = NewErrorGen("value is not greater or equal than %v")

	// ErrNumberLT when value is not smaller than
	ErrNumberLT ErrorGen = NewErrorGen("value is not smaller than %v")

	// ErrNumberLTE when value is not smaller or equal than
	ErrNumberLTE ErrorGen = NewErrorGen("value is not smaller or equal than %v")
)
