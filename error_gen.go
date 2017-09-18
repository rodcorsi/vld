package vld

import "fmt"

type ErrorGen interface {
	Gen(args ...interface{}) error
}

type errorGen struct {
	message string
}

func NewErrorGen(message string) *errorGen {
	return &errorGen{message}
}

func (e *errorGen) Gen(args ...interface{}) error {
	return fmt.Errorf(e.message, args...)
}
