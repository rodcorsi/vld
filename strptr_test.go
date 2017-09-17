package vld

import "testing"

func Test_StrPtr_Required(t *testing.T) {
	validate := New()
	ok := validate.Init("").StrPtr(nil).Ok()
	if !ok {
		t.Errorf("StrPtr.Required() = %v, want %v", ok, true)
	}
	ok = validate.Init("").StrPtr(nil).Required().Ok()
	if ok {
		t.Errorf("StrPtr.Required() = %v, want %v", ok, false)
	}
	value := ""
	ok = validate.Init("").StrPtr(&value).Required().Ok()
	if !ok {
		t.Errorf("StrPtr.Required() = %v, want %v", ok, true)
	}
}
