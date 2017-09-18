package vld

import "testing"

func Test_StrPtr_Required(t *testing.T) {
	var validate Validate
	ok := validate.Ok("", StrPtr(nil).Error())
	if !ok {
		t.Errorf("StrPtr.Required() = %v, want %v", ok, true)
	}
	ok = validate.Ok("", StrPtr(nil).Required().Error())
	if ok {
		t.Errorf("StrPtr.Required() = %v, want %v", ok, false)
	}
	value := ""
	ok = validate.Ok("", StrPtr(&value).Required().Error())
	if !ok {
		t.Errorf("StrPtr.Required() = %v, want %v", ok, true)
	}
}
