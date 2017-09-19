package vld

import (
	"errors"
	"testing"
)

func TestValidate_Ok(t *testing.T) {
	validate := New()
	if !validate.Ok("", nil) {
		t.Error("Ok() expected ok for nil error")
	}
	if validate.Ok("", errors.New("")) {
		t.Error("Ok() expected false for error")
	}
}

func TestValidate_Error(t *testing.T) {
	validate := New()
	if err := validate.Error(); err != nil {
		t.Error("Error() expected nil", err)
	}
	validate.Ok("", nil)
	if err := validate.Error(); err != nil {
		t.Error("Error() expected nil", err)
	}
	validate.Ok("", errors.New(""))
	if err := validate.Error(); err == nil {
		t.Error("Error() expected not nil")
	}
	validate.Ok("", errors.New(""))
	if err := validate.Error(); err == nil {
		t.Error("Error() expected not nil 2 errors")
	}
}
