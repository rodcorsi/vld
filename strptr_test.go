package vld

import (
	"regexp"
	"testing"
)

func ptrStr(val string) *string {
	return &val
}

func Test_StrPtr_Required(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		wantErr bool
	}{
		{"1", StrPtr(nil), true},
		{"2", StrPtr(ptrStr("x")), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Required().Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.Required() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_GT(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", StrPtr(ptrStr("a")), 0, false},    // test zero length
		{"2", StrPtr(nil), 10, false},           // not required
		{"3", StrPtr(nil).Required(), 10, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, true},     // greater than 1
		{"5", StrPtr(ptrStr("ab")), 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GT(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.GT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_GTE(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", StrPtr(ptrStr("a")), 0, false},    // test zero length
		{"2", StrPtr(nil), 10, false},           // not required
		{"3", StrPtr(nil).Required(), 10, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, false},    // equal 1
		{"5", StrPtr(ptrStr("ab")), 1, false},
		{"6", StrPtr(ptrStr("ab")), 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GTE(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.GTE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_LT(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", StrPtr(ptrStr("a")), 0, true},    // test zero length
		{"2", StrPtr(nil), 0, false},           // not required
		{"3", StrPtr(nil).Required(), 0, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, true},    // equal
		{"5", StrPtr(ptrStr("a")), 2, false},   // smaller than 2
		{"6", StrPtr(ptrStr("ab")), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LT(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.LT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_LTE(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", StrPtr(ptrStr("a")), 0, true},    // test zero length
		{"2", StrPtr(nil), 0, false},           // not required
		{"3", StrPtr(nil).Required(), 0, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, false},   // equal
		{"5", StrPtr(ptrStr("a")), 2, false},   // smaller than 2
		{"6", StrPtr(ptrStr("ab")), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LTE(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.LTE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_Length(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		min     int
		max     int
		wantErr bool
	}{
		{"1", StrPtr(ptrStr("a")), 0, 0, true},     // test zero
		{"2", StrPtr(nil), 5, 10, false},           // not required
		{"3", StrPtr(nil).Required(), 5, 10, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, 3, false},    // equal min
		{"5", StrPtr(ptrStr("ab")), 1, 3, false},   // mid
		{"6", StrPtr(ptrStr("abc")), 1, 3, false},  // equal min
		{"7", StrPtr(ptrStr("abc")), 1, 2, true},   // upper
		{"8", StrPtr(ptrStr("abc")), 4, 5, true},   // smaller
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Length(tt.min, tt.max).Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.Length() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_Len(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", StrPtr(ptrStr("a")), 0, true},    // test zero
		{"2", StrPtr(nil), 5, false},           // not required
		{"3", StrPtr(nil).Required(), 5, true}, // required
		{"4", StrPtr(ptrStr("ab")), 2, false},  // equal
		{"5", StrPtr(ptrStr("ab")), 3, true},   // different
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Len(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.Len() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_Match(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		pattern string
		wantErr bool
	}{
		{"1", StrPtr(nil), "[a-z]", false},           // not required
		{"2", StrPtr(nil).Required(), "[a-z]", true}, // required
		{"3", StrPtr(ptrStr("ab")), "[a-z]", false},  // match
		{"4", StrPtr(ptrStr("ab")), "[0-9]", true},   // not match
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Match(regexp.MustCompile(tt.pattern)).Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.Match() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_OneOf(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		values  []string
		wantErr bool
	}{
		{"1", StrPtr(nil), []string{"ab"}, false},               // not required
		{"2", StrPtr(nil).Required(), []string{"ab"}, true},     // required
		{"3", StrPtr(ptrStr("ab")), []string{"ab"}, false},      // one
		{"4", StrPtr(ptrStr("ab")), []string{"abc"}, true},      // not match
		{"5", StrPtr(ptrStr("ab")), []string{"de", "fg"}, true}, // not match
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.OneOf(tt.values).Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.OneOf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_StrPtr_IsEmail(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		wantErr bool
	}{
		{"1", StrPtr(nil), false},                   // not required
		{"2", StrPtr(nil).Required(), true},         // required
		{"3", StrPtr(ptrStr("foo@bah.com")), false}, // match
		{"4", StrPtr(ptrStr("not email")), true},    // not match
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.IsEmail().Error(); (err != nil) != tt.wantErr {
				t.Errorf("StrPtr.IsEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
