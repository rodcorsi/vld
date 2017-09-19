package vld

import (
	"testing"
)

func Test_String_Required(t *testing.T) {
	var validate Validate
	ok := validate.Ok("", String("").Error())
	if !ok {
		t.Errorf("String.Required() = %v, want %v", ok, true)
	}
	ok = validate.Ok("", String("").Required().Error())
	if ok {
		t.Errorf("String.Required() = %v, want %v", ok, false)
	}
	ok = validate.Ok("", String("x").Required().Error())
	if !ok {
		t.Errorf("String.Required() = %v, want %v", ok, true)
	}
}
