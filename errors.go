package vld

// ErrRequired when value is required
// used in String, StrPtr, Number
var ErrRequired ErrorGen = NewErrorGen("required")

// ErrGreaterThanMax when value is greater than max
// used in String, StrPtr, Number
var ErrGreaterThanMax ErrorGen = NewErrorGen("greater than max(%v)")

// ErrSmallerThanMin when value is smaller than min
// used in String, StrPtr, Number
var ErrSmallerThanMin ErrorGen = NewErrorGen("smaller than min(%v)")

// ErrLength when value is out of length
// used in String, StrPtr
var ErrLength ErrorGen = NewErrorGen("out of length(min:%v max:%v)")

// ErrLen when value is out of len
// used in String, StrPtr
var ErrLen ErrorGen = NewErrorGen("out of len(%v)")

// ErrRange when value is out of range
// used in Number
var ErrRange ErrorGen = NewErrorGen("out of range(min:%v max:%v)")
