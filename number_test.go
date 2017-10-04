package vld

import (
	"testing"
)

type testNumber struct {
	name    string
	s       *numberVld
	wantErr bool
}

func Test_Number_Required(t *testing.T) {
	testNumberRequired(t, "Number", []testNumber{
		{"1", Number(0), true},
		{"2", Number(1), false},
	})
	testNumberRequired(t, "NumF32", []testNumber{
		{"1", NumF32(0), true},
		{"2", NumF32(1), false},
	})
	testNumberRequired(t, "NumF64", []testNumber{
		{"1", NumF64(0), true},
		{"2", NumF64(1), false},
	})
	testNumberRequired(t, "NumI32", []testNumber{
		{"1", NumI32(0), true},
		{"2", NumI32(1), false},
	})
	testNumberRequired(t, "NumI64", []testNumber{
		{"1", NumI64(0), true},
		{"2", NumI64(1), false},
	})
}

func testNumberRequired(t *testing.T, detail string, tests []testNumber) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Required().Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.Required() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

type testNumberValue struct {
	name    string
	s       *numberVld
	value   float64
	wantErr bool
}

func Test_Number_GT(t *testing.T) {
	testNumberGT(t, "Number", []testNumberValue{
		{"1", Number(0), 0, false},            // test zero
		{"2", Number(0), 10, false},           // not required
		{"3", Number(0).Required(), 10, true}, // required
		{"4", Number(1), 1, true},             // greater than 1
		{"5", Number(2), 1, false},
		{"6", Number(1), 10, true},
	})
	testNumberGT(t, "NumF32", []testNumberValue{
		{"1", NumF32(0), 0, false},            // test zero
		{"2", NumF32(0), 10, false},           // not required
		{"3", NumF32(0).Required(), 10, true}, // required
		{"4", NumF32(1), 1, true},             // greater than 1
		{"5", NumF32(2), 1, false},
		{"6", NumF32(1), 10, true},
	})
	testNumberGT(t, "NumF64", []testNumberValue{
		{"1", NumF64(0), 0, false},            // test zero
		{"2", NumF64(0), 10, false},           // not required
		{"3", NumF64(0).Required(), 10, true}, // required
		{"4", NumF64(1), 1, true},             // greater than 1
		{"5", NumF64(2), 1, false},
		{"6", NumF64(1), 10, true},
	})
	testNumberGT(t, "NumI32", []testNumberValue{
		{"1", NumI32(0), 0, false},            // test zero
		{"2", NumI32(0), 10, false},           // not required
		{"3", NumI32(0).Required(), 10, true}, // required
		{"4", NumI32(1), 1, true},             // greater than 1
		{"5", NumI32(2), 1, false},
		{"6", NumI32(1), 10, true},
	})
	testNumberGT(t, "NumI64", []testNumberValue{
		{"1", NumI64(0), 0, false},            // test zero
		{"2", NumI64(0), 10, false},           // not required
		{"3", NumI64(0).Required(), 10, true}, // required
		{"4", NumI64(1), 1, true},             // greater than 1
		{"5", NumI64(2), 1, false},
		{"6", NumI64(1), 10, true},
	})
}

func testNumberGT(t *testing.T, detail string, tests []testNumberValue) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GT(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.GT() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_Number_GTE(t *testing.T) {
	testNumberGTE(t, "Number", []testNumberValue{
		{"1", Number(0), 0, false},            // test zero
		{"2", Number(0), 10, false},           // not required
		{"3", Number(0).Required(), 10, true}, // required
		{"4", Number(1), 10, true},            // greater than 1
		{"5", Number(1), 1, false},
		{"6", Number(1), 10, true},
	})
	testNumberGTE(t, "NumF32", []testNumberValue{
		{"1", NumF32(0), 0, false},            // test zero
		{"2", NumF32(0), 10, false},           // not required
		{"3", NumF32(0).Required(), 10, true}, // required
		{"4", NumF32(1), 10, true},            // greater than 1
		{"5", NumF32(1), 1, false},
		{"6", NumF32(1), 10, true},
	})
	testNumberGTE(t, "NumF64", []testNumberValue{
		{"1", NumF64(0), 0, false},            // test zero
		{"2", NumF64(0), 10, false},           // not required
		{"3", NumF64(0).Required(), 10, true}, // required
		{"4", NumF64(1), 10, true},            // greater than 1
		{"5", NumF64(1), 1, false},
		{"6", NumF64(1), 10, true},
	})
	testNumberGTE(t, "NumI32", []testNumberValue{
		{"1", NumI32(0), 0, false},            // test zero
		{"2", NumI32(0), 10, false},           // not required
		{"3", NumI32(0).Required(), 10, true}, // required
		{"4", NumI32(1), 10, true},            // greater than 1
		{"5", NumI32(1), 1, false},
		{"6", NumI32(1), 10, true},
	})
	testNumberGTE(t, "NumI64", []testNumberValue{
		{"1", NumI64(0), 0, false},            // test zero
		{"2", NumI64(0), 10, false},           // not required
		{"3", NumI64(0).Required(), 10, true}, // required
		{"4", NumI64(1), 10, true},            // greater than 1
		{"5", NumI64(1), 1, false},
		{"6", NumI64(1), 10, true},
	})
}

func testNumberGTE(t *testing.T, detail string, tests []testNumberValue) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GTE(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.GTE() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_Number_LT(t *testing.T) {
	testNumberLT(t, "Number", []testNumberValue{
		{"1", Number(0), 0, false},             // test zero
		{"2", Number(0), -10, false},           // not required
		{"3", Number(0).Required(), -10, true}, // required
		{"4", Number(1), 1, true},              // smaller than 1
		{"5", Number(1), 2, false},
		{"6", Number(10), 1, true},
	})
	testNumberLT(t, "NumF32", []testNumberValue{
		{"1", NumF32(0), 0, false},             // test zero
		{"2", NumF32(0), -10, false},           // not required
		{"3", NumF32(0).Required(), -10, true}, // required
		{"4", NumF32(1), 1, true},              // smaller than 1
		{"5", NumF32(1), 2, false},
		{"6", NumF32(10), 1, true},
	})
	testNumberLT(t, "NumF64", []testNumberValue{
		{"1", NumF64(0), 0, false},             // test zero
		{"2", NumF64(0), -10, false},           // not required
		{"3", NumF64(0).Required(), -10, true}, // required
		{"4", NumF64(1), 1, true},              // smaller than 1
		{"5", NumF64(1), 2, false},
		{"6", NumF64(10), 1, true},
	})
	testNumberLT(t, "NumI32", []testNumberValue{
		{"1", NumI32(0), 0, false},             // test zero
		{"2", NumI32(0), -10, false},           // not required
		{"3", NumI32(0).Required(), -10, true}, // required
		{"4", NumI32(1), 1, true},              // smaller than 1
		{"5", NumI32(1), 2, false},
		{"6", NumI32(10), 1, true},
	})
	testNumberLT(t, "NumI64", []testNumberValue{
		{"1", NumI64(0), 0, false},             // test zero
		{"2", NumI64(0), -10, false},           // not required
		{"3", NumI64(0).Required(), -10, true}, // required
		{"4", NumI64(1), 1, true},              // smaller than 1
		{"5", NumI64(1), 2, false},
		{"6", NumI64(10), 1, true},
	})
}

func testNumberLT(t *testing.T, detail string, tests []testNumberValue) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LT(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.LT() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_Number_LTE(t *testing.T) {
	testNumberLTE(t, "Number", []testNumberValue{
		{"1", Number(0), 0, false},             // test zero
		{"2", Number(0), -10, false},           // not required
		{"3", Number(0).Required(), -10, true}, // required
		{"4", Number(1), 1, false},             // smaller than 1
		{"5", Number(1), 2, false},
		{"6", Number(10), 1, true},
	})
	testNumberLTE(t, "NumF32", []testNumberValue{
		{"1", NumF32(0), 0, false},             // test zero
		{"2", NumF32(0), -10, false},           // not required
		{"3", NumF32(0).Required(), -10, true}, // required
		{"4", NumF32(1), 1, false},             // smaller than 1
		{"5", NumF32(1), 2, false},
		{"6", NumF32(10), 1, true},
	})
	testNumberLTE(t, "NumF64", []testNumberValue{
		{"1", NumF64(0), 0, false},             // test zero
		{"2", NumF64(0), -10, false},           // not required
		{"3", NumF64(0).Required(), -10, true}, // required
		{"4", NumF64(1), 1, false},             // smaller than 1
		{"5", NumF64(1), 2, false},
		{"6", NumF64(10), 1, true},
	})
	testNumberLTE(t, "NumI32", []testNumberValue{
		{"1", NumI32(0), 0, false},             // test zero
		{"2", NumI32(0), -10, false},           // not required
		{"3", NumI32(0).Required(), -10, true}, // required
		{"4", NumI32(1), 1, false},             // smaller than 1
		{"5", NumI32(1), 2, false},
		{"6", NumI32(10), 1, true},
	})
	testNumberLTE(t, "NumI64", []testNumberValue{
		{"1", NumI64(0), 0, false},             // test zero
		{"2", NumI64(0), -10, false},           // not required
		{"3", NumI64(0).Required(), -10, true}, // required
		{"4", NumI64(1), 1, false},             // smaller than 1
		{"5", NumI64(1), 2, false},
		{"6", NumI64(10), 1, true},
	})
}

func testNumberLTE(t *testing.T, detail string, tests []testNumberValue) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LTE(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.LTE() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

type testNumberMinMax struct {
	name    string
	s       *numberVld
	min     float64
	max     float64
	wantErr bool
}

func Test_Number_Range(t *testing.T) {
	testNumberRange(t, "Number", []testNumberMinMax{
		{"1", Number(0), 0, 0, false},           // test zero
		{"2", Number(0), 1, 2, false},           // not required
		{"3", Number(0).Required(), 1, 2, true}, // required
		{"4", Number(1), 2, 4, true},
		{"5", Number(2), 2, 4, false},
		{"6", Number(3), 2, 4, false},
		{"7", Number(4), 2, 4, false},
		{"7", Number(5), 2, 4, true},
	})
	testNumberRange(t, "NumF32", []testNumberMinMax{
		{"1", NumF32(0), 0, 0, false},           // test zero
		{"2", NumF32(0), 1, 2, false},           // not required
		{"3", NumF32(0).Required(), 1, 2, true}, // required
		{"4", NumF32(1), 2, 4, true},
		{"5", NumF32(2), 2, 4, false},
		{"6", NumF32(3), 2, 4, false},
		{"7", NumF32(4), 2, 4, false},
		{"7", NumF32(5), 2, 4, true},
	})
	testNumberRange(t, "NumF64", []testNumberMinMax{
		{"1", NumF64(0), 0, 0, false},           // test zero
		{"2", NumF64(0), 1, 2, false},           // not required
		{"3", NumF64(0).Required(), 1, 2, true}, // required
		{"4", NumF64(1), 2, 4, true},
		{"5", NumF64(2), 2, 4, false},
		{"6", NumF64(3), 2, 4, false},
		{"7", NumF64(4), 2, 4, false},
		{"7", NumF64(5), 2, 4, true},
	})
	testNumberRange(t, "NumI32", []testNumberMinMax{
		{"1", NumI32(0), 0, 0, false},           // test zero
		{"2", NumI32(0), 1, 2, false},           // not required
		{"3", NumI32(0).Required(), 1, 2, true}, // required
		{"4", NumI32(1), 2, 4, true},
		{"5", NumI32(2), 2, 4, false},
		{"6", NumI32(3), 2, 4, false},
		{"7", NumI32(4), 2, 4, false},
		{"7", NumI32(5), 2, 4, true},
	})
	testNumberRange(t, "NumI64", []testNumberMinMax{
		{"1", NumI64(0), 0, 0, false},           // test zero
		{"2", NumI64(0), 1, 2, false},           // not required
		{"3", NumI64(0).Required(), 1, 2, true}, // required
		{"4", NumI64(1), 2, 4, true},
		{"5", NumI64(2), 2, 4, false},
		{"6", NumI64(3), 2, 4, false},
		{"7", NumI64(4), 2, 4, false},
		{"7", NumI64(5), 2, 4, true},
	})
}

func testNumberRange(t *testing.T, detail string, tests []testNumberMinMax) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Range(tt.min, tt.max).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.Range() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}
