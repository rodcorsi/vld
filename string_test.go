package vld

import (
	"testing"
)

func Test_String_Required(t *testing.T) {
	validate := New()
	ok := validate.Init("").String("").Ok()
	if !ok {
		t.Errorf("String.Required() = %v, want %v", ok, true)
	}
	ok = validate.Init("").String("").Required().Ok()
	if ok {
		t.Errorf("String.Required() = %v, want %v", ok, false)
	}
	ok = validate.Init("").String("x").Required().Ok()
	if !ok {
		t.Errorf("String.Required() = %v, want %v", ok, true)
	}
}

func Test_String_Max(t *testing.T) {
	validate := New()
	ok := validate.Init("").String("").Max(0).Ok()
	if !ok {
		t.Errorf("String.Max() = %v, want %v", ok, true)
	}
	ok = validate.Init("").String("x").Max(0).Ok()
	if ok {
		t.Errorf("String.Max() = %v, want %v", ok, false)
	}
	ok = validate.Init("").String("x").Max(1).Ok()
	if !ok {
		t.Errorf("String.Max() = %v, want %v", ok, true)
	}
}
