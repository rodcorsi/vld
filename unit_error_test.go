package vld

import (
	"testing"
)

func Test_unitError_ErrorID(t *testing.T) {
	unitError := NewUnitError("id1", nil)
	if unitError.ErrorID() != "id1" {
		t.Errorf("ErrorID() = '%v' want: 'id1'", unitError.ErrorID())
	}
}

func Test_unitError_Args(t *testing.T) {
	unitError := NewUnitError("", nil)
	if unitError.Args() != nil {
		t.Errorf("Args() = '%v' want: nil", unitError.Args())
	}

	unitError = NewUnitError("", Args{})
	if args := unitError.Args(); args == nil || len(args) != 0 {
		t.Errorf("Args() = '%v' want: {}", args)
	}

	unitError = NewUnitError("", Args{"a", 1})
	if args := unitError.Args(); len(args) != 2 || args[0] != "a" || args[1] != 1 {
		t.Errorf("Args() = '%v' want: {'a', 1}", args)
	}
}

func Test_unitError_Error(t *testing.T) {
	tests := []struct {
		name string
		u    *unitError
		want string
	}{
		{"1", &unitError{"", nil}, "[]"},
		{"2", &unitError{"id", nil}, "id[]"},
		{"3", &unitError{"id", Args{"a"}}, "id[a]"},
		{"4", &unitError{"id", Args{"a", "b"}}, "id[a b]"},
		{"5", &unitError{ErrRequired, nil}, defaultErrMessage[ErrRequired]},
		{"6", &unitError{ErrValidation, Args{"f1", "foo"}}, "validation for 'f1' failed: foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Error(); got != tt.want {
				t.Errorf("unitError.Error() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
