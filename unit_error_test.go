package vld

import "testing"

func Test_unitError_ErrorID(t *testing.T) {
	unitError := newUnitError("id1", nil)
	if unitError.ErrorID() != "id1" {
		t.Errorf("ErrorID() = '%v' want: 'id1'", unitError.ErrorID())
	}
}

func Test_unitError_Args(t *testing.T) {
	unitError := newUnitError("", nil)
	if unitError.Args() != nil {
		t.Errorf("Args() = '%v' want: nil", unitError.Args())
	}

	unitError = newUnitError("", Args{})
	if args := unitError.Args(); args == nil || len(args) != 0 {
		t.Errorf("Args() = '%' want: {}", args)
	}

	unitError = newUnitError("", Args{"a", 1})
	if args := unitError.Args(); len(args) != 2 || args[0] != "a" || args[1] != 1 {
		t.Errorf("Args() = '%' want: {'a', 1}", args)
	}
}
