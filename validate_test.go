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

func TestValidate_HasError(t *testing.T) {
	validate := New()
	if validate.HasError() {
		t.Error("HasError() expected false")
	}
	validate.Ok("", errors.New(""))
	if !validate.HasError() {
		t.Error("HasError() expected true")
	}
	validate.Ok("", errors.New(""))
	if !validate.HasError() {
		t.Error("HasError() expected true 2 errors")
	}
}
