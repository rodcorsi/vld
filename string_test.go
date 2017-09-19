package vld

import (
	"regexp"
	"testing"
)

func Test_String_Required(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		wantErr bool
	}{
		{"1", String(""), true},
		{"2", String("x"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Required().Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.Required() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_GT(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", String("a"), 0, false},           // test zero length
		{"2", String(""), 10, false},           // not required
		{"3", String("").Required(), 10, true}, // required
		{"4", String("a"), 1, true},            // greater than 1
		{"5", String("ab"), 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GT(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.GT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_GTE(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", String("a"), 0, false},           // test zero length
		{"2", String(""), 10, false},           // not required
		{"3", String("").Required(), 10, true}, // required
		{"4", String("a"), 1, false},           // equal 1
		{"5", String("ab"), 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GTE(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.GTE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_LT(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", String("a"), 0, true},           // test zero length
		{"2", String(""), 0, false},           // not required
		{"3", String("").Required(), 0, true}, // required
		{"4", String("a"), 1, true},           // equal
		{"5", String("a"), 2, false},          // smaller than 2
		{"6", String("ab"), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LT(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.LT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_LTE(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", String("a"), 0, true},           // test zero length
		{"2", String(""), 0, false},           // not required
		{"3", String("").Required(), 0, true}, // required
		{"4", String("a"), 1, false},          // equal
		{"5", String("a"), 2, false},          // smaller than 2
		{"6", String("ab"), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LTE(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.LTE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_Length(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		min     int
		max     int
		wantErr bool
	}{
		{"1", String("a"), 0, 0, true},            // test zero
		{"2", String(""), 5, 10, false},           // not required
		{"3", String("").Required(), 5, 10, true}, // required
		{"4", String("a"), 1, 3, false},           // equal min
		{"5", String("ab"), 1, 3, false},          // mid
		{"6", String("abc"), 1, 3, false},         // equal min
		{"7", String("abc"), 1, 2, true},          // upper
		{"8", String("abc"), 4, 5, true},          // smaller
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Length(tt.min, tt.max).Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.Length() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_Len(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		length  int
		wantErr bool
	}{
		{"1", String("a"), 0, true},           // test zero
		{"2", String(""), 5, false},           // not required
		{"3", String("").Required(), 5, true}, // required
		{"4", String("ab"), 2, false},         // equal
		{"5", String("ab"), 3, true},          // different
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Len(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.Len() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_Match(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		pattern string
		wantErr bool
	}{
		{"1", String(""), "[a-z]", false},           // not required
		{"2", String("").Required(), "[a-z]", true}, // required
		{"3", String("ab"), "[a-z]", false},         // match
		{"4", String("ab"), "[0-9]", true},          // not match
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Match(regexp.MustCompile(tt.pattern)).Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.Match() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_OneOf(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		values  []string
		wantErr bool
	}{
		{"1", String(""), []string{"ab"}, false},           // not required
		{"2", String("").Required(), []string{"ab"}, true}, // required
		{"3", String("ab"), []string{"ab"}, false},         // one
		{"4", String("ab"), []string{"abc"}, true},         // not match
		{"5", String("ab"), []string{"de", "fg"}, true},    // not match
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.OneOf(tt.values).Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.OneOf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_String_IsEmail(t *testing.T) {
	tests := []struct {
		name    string
		s       *stringVld
		wantErr bool
	}{
		{"1", String(""), false},            // not required
		{"2", String("").Required(), true},  // required
		{"3", String("foo@bah.com"), false}, // match
		{"4", String("not email"), true},    // not match
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.IsEmail().Error(); (err != nil) != tt.wantErr {
				t.Errorf("String.IsEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
