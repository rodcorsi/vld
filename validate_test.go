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

func Test_Errors_Error(t *testing.T) {
	tests := []struct {
		name string
		e    Errors
		want string
	}{
		{"1", Errors{}, ""},
		{"2", Errors{errors.New("a")}, "a"},
		{"3", Errors{errors.New("a"), errors.New("b")}, "a\nb"},
		{"4", Errors{errors.New("a"), errors.New("b"), errors.New("c")}, "a\nb\nc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("Errors.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
