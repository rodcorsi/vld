package vld

// ErrRequired when value is required
// used in String, StrPtr, Number
var ErrRequired ErrorGen = NewErrorGen("required")

// ErrStringGT when length is not greater than
// used in String, StrPtr
var ErrStringGT ErrorGen = NewErrorGen("length is not greater than %v")

// ErrStringGTE when length is not greater or equal than
// used in String, StrPtr
var ErrStringGTE ErrorGen = NewErrorGen("length is not greater or equal than %v")

// ErrStringLT when length is not smaller than
// used in String, StrPtr
var ErrStringLT ErrorGen = NewErrorGen("length is not smaller than %v")

// ErrStringLTE when length is not smaller or equal than
// used in String, StrPtr
var ErrStringLTE ErrorGen = NewErrorGen("length is not smaller or equal than %v")

// ErrLength when value is out of value
// used in String, StrPtr
var ErrLength ErrorGen = NewErrorGen("out of value(min:%v max:%v)")

// ErrLen when value is out of len
// used in String, StrPtr
var ErrLen ErrorGen = NewErrorGen("out of len(%v)")

// ErrRange when value is out of range
// used in Number
var ErrRange ErrorGen = NewErrorGen("out of range(min:%v max:%v)")

// ErrNumberGT when value is not greater than
// used in Number
var ErrNumberGT ErrorGen = NewErrorGen("value is not greater than %v")

// ErrNumberGTE when value is not greater or equal than
// used in Number
var ErrNumberGTE ErrorGen = NewErrorGen("value is not greater or equal than %v")

// ErrNumberLT when value is not smaller than
// used in Number
var ErrNumberLT ErrorGen = NewErrorGen("value is not smaller than %v")

// ErrNumberLTE when value is not smaller or equal than
// used in Number
var ErrNumberLTE ErrorGen = NewErrorGen("value is not smaller or equal than %v")
