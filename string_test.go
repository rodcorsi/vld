package vld

import (
	"regexp"
	"testing"
)

func ptrStr(val string) *string {
	return &val
}

type testString struct {
	name    string
	s       *stringVld
	wantErr bool
}

func Test_String_Required(t *testing.T) {
	testStringRequired(t, "String", []testString{
		{"1", String(""), true},
		{"2", String("x"), false},
	})
	testStringRequired(t, "StrPtr", []testString{
		{"1", StrPtr(nil), true},
		{"2", StrPtr(ptrStr("x")), false},
	})
}

func testStringRequired(t *testing.T, detail string, tests []testString) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Required().Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.Required() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

type testStringLength struct {
	name    string
	s       *stringVld
	length  int
	wantErr bool
}

func Test_String_GT(t *testing.T) {
	testStringGT(t, "String", []testStringLength{
		{"1", String("a"), 0, false},           // test zero length
		{"2", String(""), 10, false},           // not required
		{"3", String("").Required(), 10, true}, // required
		{"4", String("a"), 1, true},            // greater than 1
		{"5", String("ab"), 1, false},
	})

	testStringGT(t, "StrPtr", []testStringLength{
		{"1", StrPtr(ptrStr("a")), 0, false},    // test zero length
		{"2", StrPtr(nil), 10, false},           // not required
		{"3", StrPtr(nil).Required(), 10, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, true},     // greater than 1
		{"5", StrPtr(ptrStr("ab")), 1, false},
	})
}

func testStringGT(t *testing.T, detail string, tests []testStringLength) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GT(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.GT() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_String_GTE(t *testing.T) {
	testStringGTE(t, "String", []testStringLength{
		{"1", String("a"), 0, false},           // test zero length
		{"2", String(""), 10, false},           // not required
		{"3", String("").Required(), 10, true}, // required
		{"4", String("a"), 1, false},           // equal 1
		{"5", String("ab"), 1, false},
		{"6", String("ab"), 10, true},
	})

	testStringGTE(t, "StrPtr", []testStringLength{
		{"1", StrPtr(ptrStr("a")), 0, false},    // test zero length
		{"2", StrPtr(nil), 10, false},           // not required
		{"3", StrPtr(nil).Required(), 10, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, false},    // equal 1
		{"5", StrPtr(ptrStr("ab")), 1, false},
		{"6", StrPtr(ptrStr("ab")), 10, true},
	})
}

func testStringGTE(t *testing.T, detail string, tests []testStringLength) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GTE(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.GTE() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_String_LT(t *testing.T) {
	testStringLT(t, "String", []testStringLength{
		{"1", String("a"), 0, true},           // test zero length
		{"2", String(""), 0, false},           // not required
		{"3", String("").Required(), 0, true}, // required
		{"4", String("a"), 1, true},           // equal
		{"5", String("a"), 2, false},          // smaller than 2
		{"6", String("ab"), 1, true},
	})

	testStringLT(t, "StrPtr", []testStringLength{
		{"1", StrPtr(ptrStr("a")), 0, true},    // test zero length
		{"2", StrPtr(nil), 0, false},           // not required
		{"3", StrPtr(nil).Required(), 0, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, true},    // equal
		{"5", StrPtr(ptrStr("a")), 2, false},   // smaller than 2
		{"6", StrPtr(ptrStr("ab")), 1, true},
	})
}

func testStringLT(t *testing.T, detail string, tests []testStringLength) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LT(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.LT() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_String_LTE(t *testing.T) {
	testStringLTE(t, "String", []testStringLength{
		{"1", String("a"), 0, true},           // test zero length
		{"2", String(""), 0, false},           // not required
		{"3", String("").Required(), 0, true}, // required
		{"4", String("a"), 1, false},          // equal
		{"5", String("a"), 2, false},          // smaller than 2
		{"6", String("ab"), 1, true},
	})

	testStringLTE(t, "StrPtr", []testStringLength{
		{"1", StrPtr(ptrStr("a")), 0, true},    // test zero length
		{"2", StrPtr(nil), 0, false},           // not required
		{"3", StrPtr(nil).Required(), 0, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, false},   // equal
		{"5", StrPtr(ptrStr("a")), 2, false},   // smaller than 2
		{"6", StrPtr(ptrStr("ab")), 1, true},
	})
}

func testStringLTE(t *testing.T, detail string, tests []testStringLength) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LTE(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.LTE() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

type testStringMinMax struct {
	name    string
	s       *stringVld
	min     int
	max     int
	wantErr bool
}

func Test_String_Length(t *testing.T) {
	testStringLength(t, "String", []testStringMinMax{
		{"1", String("a"), 0, 0, true},            // test zero
		{"2", String(""), 5, 10, false},           // not required
		{"3", String("").Required(), 5, 10, true}, // required
		{"4", String("a"), 1, 3, false},           // equal min
		{"5", String("ab"), 1, 3, false},          // mid
		{"6", String("abc"), 1, 3, false},         // equal min
		{"7", String("abc"), 1, 2, true},          // upper
		{"8", String("abc"), 4, 5, true},          // smaller
	})

	testStringLength(t, "StrPtr", []testStringMinMax{
		{"1", StrPtr(ptrStr("a")), 0, 0, true},     // test zero
		{"2", StrPtr(nil), 5, 10, false},           // not required
		{"3", StrPtr(nil).Required(), 5, 10, true}, // required
		{"4", StrPtr(ptrStr("a")), 1, 3, false},    // equal min
		{"5", StrPtr(ptrStr("ab")), 1, 3, false},   // mid
		{"6", StrPtr(ptrStr("abc")), 1, 3, false},  // equal min
		{"7", StrPtr(ptrStr("abc")), 1, 2, true},   // upper
		{"8", StrPtr(ptrStr("abc")), 4, 5, true},   // smaller
	})
}

func testStringLength(t *testing.T, detail string, tests []testStringMinMax) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Length(tt.min, tt.max).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.Length() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_String_Len(t *testing.T) {
	testStringLen(t, "String", []testStringLength{
		{"1", String("a"), 0, true},           // test zero
		{"2", String(""), 5, false},           // not required
		{"3", String("").Required(), 5, true}, // required
		{"4", String("ab"), 2, false},         // equal
		{"5", String("ab"), 3, true},          // different
	})
	testStringLen(t, "StrPtr", []testStringLength{
		{"1", StrPtr(ptrStr("a")), 0, true},    // test zero
		{"2", StrPtr(nil), 5, false},           // not required
		{"3", StrPtr(nil).Required(), 5, true}, // required
		{"4", StrPtr(ptrStr("ab")), 2, false},  // equal
		{"5", StrPtr(ptrStr("ab")), 3, true},   // different
	})
}

func testStringLen(t *testing.T, detail string, tests []testStringLength) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Len(tt.length).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.Len() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

type testStringMatch struct {
	name    string
	s       *stringVld
	pattern string
	wantErr bool
}

func Test_String_Match(t *testing.T) {
	testStringMatch(t, "String", []testStringMatch{
		{"1", String(""), "[a-z]", false},           // not required
		{"2", String("").Required(), "[a-z]", true}, // required
		{"3", String("ab"), "[a-z]", false},         // match
		{"4", String("ab"), "[0-9]", true},          // not match
	})
	testStringMatch(t, "StrPtr", []testStringMatch{
		{"1", StrPtr(nil), "[a-z]", false},           // not required
		{"2", StrPtr(nil).Required(), "[a-z]", true}, // required
		{"3", StrPtr(ptrStr("ab")), "[a-z]", false},  // match
		{"4", StrPtr(ptrStr("ab")), "[0-9]", true},   // not match
	})
}

func testStringMatch(t *testing.T, detail string, tests []testStringMatch) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Match(regexp.MustCompile(tt.pattern)).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.Match() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

type testStringOneOf struct {
	name    string
	s       *stringVld
	values  []string
	wantErr bool
}

func Test_String_OneOf(t *testing.T) {
	testStringOneOf(t, "String", []testStringOneOf{
		{"1", String(""), []string{"ab"}, false},           // not required
		{"2", String("").Required(), []string{"ab"}, true}, // required
		{"3", String("ab"), []string{"ab"}, false},         // one
		{"4", String("ab"), []string{"abc"}, true},         // not match
		{"5", String("ab"), []string{"de", "fg"}, true},    // not match
	})
	testStringOneOf(t, "StrPtr", []testStringOneOf{
		{"1", StrPtr(nil), []string{"ab"}, false},               // not required
		{"2", StrPtr(nil).Required(), []string{"ab"}, true},     // required
		{"3", StrPtr(ptrStr("ab")), []string{"ab"}, false},      // one
		{"4", StrPtr(ptrStr("ab")), []string{"abc"}, true},      // not match
		{"5", StrPtr(ptrStr("ab")), []string{"de", "fg"}, true}, // not match
	})
}

func testStringOneOf(t *testing.T, detail string, tests []testStringOneOf) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.OneOf(tt.values).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.OneOf() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_String_IsEmail(t *testing.T) {
	testStringIsEmail(t, "String", []testString{
		{"1", String(""), false},            // not required
		{"2", String("").Required(), true},  // required
		{"3", String("foo@bah.com"), false}, // match
		{"4", String("not email"), true},    // not match
	})
	testStringIsEmail(t, "StrPtr", []testString{
		{"1", StrPtr(nil), false},                   // not required
		{"2", StrPtr(nil).Required(), true},         // required
		{"3", StrPtr(ptrStr("foo@bah.com")), false}, // match
		{"4", StrPtr(ptrStr("not email")), true},    // not match
	})
}

func testStringIsEmail(t *testing.T, detail string, tests []testString) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.IsEmail().Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.IsEmail() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}
