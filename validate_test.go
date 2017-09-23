package vld

import (
	"testing"
)

func TestValidate_Ok(t *testing.T) {
	validate := New()
	if !validate.Ok("", nil) {
		t.Error("Ok() expected ok for nil error")
	}
	if validate.Ok("", NewUnitError("", nil)) {
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
	validate.Ok("", NewUnitError("", nil))
	if err := validate.Error(); err == nil {
		t.Error("Error() expected not nil")
	}
	validate.Ok("", NewUnitError("", nil))
	if err := validate.Error(); err == nil {
		t.Error("Error() expected not nil 2 errors")
	}
}

func TestValidate_FieldError(t *testing.T) {
	validate := New()

	validate.Ok("", nil)
	fr := validate.FieldError()
	if fr != nil {
		t.Error("FieldError() expected nil")
	}

	validate.Ok("field1", NewUnitError("errorID1", nil))
	fr = validate.FieldError()
	if len(fr) != 1 {
		t.Errorf("FieldError() expected len == 1 result:%v", fr)
	}
	if fr[0].Field() != "field1" || fr[0].UnitError().ErrorID() != "errorID1" {
		t.Errorf("FieldError() expected {field1, errorID1} result:%+v", fr)
	}

	validate.Ok("field2", NewUnitError("errorID2", nil))
	fr = validate.FieldError()
	if len(fr) != 2 {
		t.Errorf("FieldError() expected len == 2 result:%v", fr)
	}
	if fr[0].Field() != "field1" || fr[0].UnitError().ErrorID() != "errorID1" || fr[1].Field() != "field2" || fr[1].UnitError().ErrorID() != "errorID2" {
		t.Errorf("FieldError() expected {{field1, errorID1},{field2, errorID2}} result:%+v", fr)
	}
}
