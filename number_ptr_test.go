package vld

import (
	"testing"
)

func ptrNum(val int) *int         { return &val }
func ptrI32(val int32) *int32     { return &val }
func ptrI64(val int64) *int64     { return &val }
func ptrF64(val float64) *float64 { return &val }
func ptrF32(val float32) *float32 { return &val }

func Test_NumberPtr_Required(t *testing.T) {
	testNumberPtrRequired(t, "NumberPtr", []testNumber{
		{"1", NumberPtr(nil), true},
		{"2", NumberPtr(ptrNum(1)), false},
	})
	testNumberPtrRequired(t, "NumF32Ptr", []testNumber{
		{"1", NumF32Ptr(nil), true},
		{"2", NumF32Ptr(ptrF32(1)), false},
	})
	testNumberPtrRequired(t, "NumF64Ptr", []testNumber{
		{"1", NumF64Ptr(nil), true},
		{"2", NumF64Ptr(ptrF64(1)), false},
	})
	testNumberPtrRequired(t, "NumI32Ptr", []testNumber{
		{"1", NumI32Ptr(nil), true},
		{"2", NumI32Ptr(ptrI32(1)), false},
	})
	testNumberPtrRequired(t, "NumI64Ptr", []testNumber{
		{"1", NumI64Ptr(nil), true},
		{"2", NumI64Ptr(ptrI64(1)), false},
	})
}

func testNumberPtrRequired(t *testing.T, detail string, tests []testNumber) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Required().Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.Required() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_NumberPtr_GT(t *testing.T) {
	testNumberPtrGT(t, "NumberPtr", []testNumberValue{
		{"1", NumberPtr(ptrNum(0)), 0, true},       // test zero
		{"2", NumberPtr(nil), 10, false},           // not required
		{"3", NumberPtr(nil).Required(), 10, true}, // required
		{"4", NumberPtr(ptrNum(1)), 1, true},       // greater than 1
		{"5", NumberPtr(ptrNum(2)), 1, false},
		{"6", NumberPtr(ptrNum(1)), 10, true},
	})
	testNumberPtrGT(t, "NumF32Ptr", []testNumberValue{
		{"1", NumF32Ptr(ptrF32(0)), 0, true},       // test zero
		{"2", NumF32Ptr(nil), 10, false},           // not required
		{"3", NumF32Ptr(nil).Required(), 10, true}, // required
		{"4", NumF32Ptr(ptrF32(1)), 1, true},       // greater than 1
		{"5", NumF32Ptr(ptrF32(2)), 1, false},
		{"6", NumF32Ptr(ptrF32(1)), 10, true},
	})
	testNumberPtrGT(t, "NumF64Ptr", []testNumberValue{
		{"1", NumF64Ptr(ptrF64(0)), 0, true},       // test zero
		{"2", NumF64Ptr(nil), 10, false},           // not required
		{"3", NumF64Ptr(nil).Required(), 10, true}, // required
		{"4", NumF64Ptr(ptrF64(1)), 1, true},       // greater than 1
		{"5", NumF64Ptr(ptrF64(2)), 1, false},
		{"6", NumF64Ptr(ptrF64(1)), 10, true},
	})
	testNumberPtrGT(t, "NumI32Ptr", []testNumberValue{
		{"1", NumI32Ptr(ptrI32(0)), 0, true},       // test zero
		{"2", NumI32Ptr(nil), 10, false},           // not required
		{"3", NumI32Ptr(nil).Required(), 10, true}, // required
		{"4", NumI32Ptr(ptrI32(1)), 1, true},       // greater than 1
		{"5", NumI32Ptr(ptrI32(2)), 1, false},
		{"6", NumI32Ptr(ptrI32(1)), 10, true},
	})
	testNumberPtrGT(t, "NumI64Ptr", []testNumberValue{
		{"1", NumI64Ptr(ptrI64(0)), 0, true},       // test zero
		{"2", NumI64Ptr(nil), 10, false},           // not required
		{"3", NumI64Ptr(nil).Required(), 10, true}, // required
		{"4", NumI64Ptr(ptrI64(1)), 1, true},       // greater than 1
		{"5", NumI64Ptr(ptrI64(2)), 1, false},
		{"6", NumI64Ptr(ptrI64(1)), 10, true},
	})
}

func testNumberPtrGT(t *testing.T, detail string, tests []testNumberValue) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GT(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.GT() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_NumberPtr_GTE(t *testing.T) {
	testNumberPtrGTE(t, "NumberPtr", []testNumberValue{
		{"1", NumberPtr(ptrNum(0)), 0, false},      // test zero
		{"2", NumberPtr(nil), 10, false},           // not required
		{"3", NumberPtr(nil).Required(), 10, true}, // required
		{"4", NumberPtr(ptrNum(1)), 10, true},      // greater than 1
		{"5", NumberPtr(ptrNum(1)), 1, false},
		{"6", NumberPtr(ptrNum(1)), 10, true},
	})
	testNumberPtrGTE(t, "NumF32Ptr", []testNumberValue{
		{"1", NumF32Ptr(ptrF32(0)), 0, false},      // test zero
		{"2", NumF32Ptr(nil), 10, false},           // not required
		{"3", NumF32Ptr(nil).Required(), 10, true}, // required
		{"4", NumF32Ptr(ptrF32(1)), 10, true},      // greater than 1
		{"5", NumF32Ptr(ptrF32(1)), 1, false},
		{"6", NumF32Ptr(ptrF32(1)), 10, true},
	})
	testNumberPtrGTE(t, "NumF64Ptr", []testNumberValue{
		{"1", NumF64Ptr(ptrF64(0)), 0, false},      // test zero
		{"2", NumF64Ptr(nil), 10, false},           // not required
		{"3", NumF64Ptr(nil).Required(), 10, true}, // required
		{"4", NumF64Ptr(ptrF64(1)), 10, true},      // greater than 1
		{"5", NumF64Ptr(ptrF64(1)), 1, false},
		{"6", NumF64Ptr(ptrF64(1)), 10, true},
	})
	testNumberPtrGTE(t, "NumI32Ptr", []testNumberValue{
		{"1", NumI32Ptr(ptrI32(0)), 0, false},      // test zero
		{"2", NumI32Ptr(nil), 10, false},           // not required
		{"3", NumI32Ptr(nil).Required(), 10, true}, // required
		{"4", NumI32Ptr(ptrI32(1)), 10, true},      // greater than 1
		{"5", NumI32Ptr(ptrI32(1)), 1, false},
		{"6", NumI32Ptr(ptrI32(1)), 10, true},
	})
	testNumberPtrGTE(t, "NumI64Ptr", []testNumberValue{
		{"1", NumI64Ptr(ptrI64(0)), 0, false},      // test zero
		{"2", NumI64Ptr(nil), 10, false},           // not required
		{"3", NumI64Ptr(nil).Required(), 10, true}, // required
		{"4", NumI64Ptr(ptrI64(1)), 10, true},      // greater than 1
		{"5", NumI64Ptr(ptrI64(1)), 1, false},
		{"6", NumI64Ptr(ptrI64(1)), 10, true},
	})
}

func testNumberPtrGTE(t *testing.T, detail string, tests []testNumberValue) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GTE(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.GTE() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_NumberPtr_LT(t *testing.T) {
	testNumberPtrLT(t, "NumberPtr", []testNumberValue{
		{"1", NumberPtr(ptrNum(0)), 0, true},        // test zero
		{"2", NumberPtr(nil), -10, false},           // not required
		{"3", NumberPtr(nil).Required(), -10, true}, // required
		{"4", NumberPtr(ptrNum(1)), 1, true},        // smaller than 1
		{"5", NumberPtr(ptrNum(1)), 2, false},
		{"6", NumberPtr(ptrNum(10)), 1, true},
	})
	testNumberPtrLT(t, "NumF32Ptr", []testNumberValue{
		{"1", NumF32Ptr(ptrF32(0)), 0, true},        // test zero
		{"2", NumF32Ptr(nil), -10, false},           // not required
		{"3", NumF32Ptr(nil).Required(), -10, true}, // required
		{"4", NumF32Ptr(ptrF32(1)), 1, true},        // smaller than 1
		{"5", NumF32Ptr(ptrF32(1)), 2, false},
		{"6", NumF32Ptr(ptrF32(10)), 1, true},
	})
	testNumberPtrLT(t, "NumF64Ptr", []testNumberValue{
		{"1", NumF64Ptr(ptrF64(0)), 0, true},        // test zero
		{"2", NumF64Ptr(nil), -10, false},           // not required
		{"3", NumF64Ptr(nil).Required(), -10, true}, // required
		{"4", NumF64Ptr(ptrF64(1)), 1, true},        // smaller than 1
		{"5", NumF64Ptr(ptrF64(1)), 2, false},
		{"6", NumF64Ptr(ptrF64(10)), 1, true},
	})
	testNumberPtrLT(t, "NumI32Ptr", []testNumberValue{
		{"1", NumI32Ptr(ptrI32(0)), 0, true},        // test zero
		{"2", NumI32Ptr(nil), -10, false},           // not required
		{"3", NumI32Ptr(nil).Required(), -10, true}, // required
		{"4", NumI32Ptr(ptrI32(1)), 1, true},        // smaller than 1
		{"5", NumI32Ptr(ptrI32(1)), 2, false},
		{"6", NumI32Ptr(ptrI32(10)), 1, true},
	})
	testNumberPtrLT(t, "NumI64Ptr", []testNumberValue{
		{"1", NumI64Ptr(ptrI64(0)), 0, true},        // test zero
		{"2", NumI64Ptr(nil), -10, false},           // not required
		{"3", NumI64Ptr(nil).Required(), -10, true}, // required
		{"4", NumI64Ptr(ptrI64(1)), 1, true},        // smaller than 1
		{"5", NumI64Ptr(ptrI64(1)), 2, false},
		{"6", NumI64Ptr(ptrI64(10)), 1, true},
	})
}

func testNumberPtrLT(t *testing.T, detail string, tests []testNumberValue) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LT(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.LT() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_NumberPtr_LTE(t *testing.T) {
	testNumberPtrLTE(t, "NumberPtr", []testNumberValue{
		{"1", NumberPtr(ptrNum(0)), 0, false},       // test zero
		{"2", NumberPtr(nil), -10, false},           // not required
		{"3", NumberPtr(nil).Required(), -10, true}, // required
		{"4", NumberPtr(ptrNum(1)), 1, false},       // smaller than 1
		{"5", NumberPtr(ptrNum(1)), 2, false},
		{"6", NumberPtr(ptrNum(10)), 1, true},
	})
	testNumberPtrLTE(t, "NumF32Ptr", []testNumberValue{
		{"1", NumF32Ptr(ptrF32(0)), 0, false},       // test zero
		{"2", NumF32Ptr(nil), -10, false},           // not required
		{"3", NumF32Ptr(nil).Required(), -10, true}, // required
		{"4", NumF32Ptr(ptrF32(1)), 1, false},       // smaller than 1
		{"5", NumF32Ptr(ptrF32(1)), 2, false},
		{"6", NumF32Ptr(ptrF32(10)), 1, true},
	})
	testNumberPtrLTE(t, "NumF64Ptr", []testNumberValue{
		{"1", NumF64Ptr(ptrF64(0)), 0, false},       // test zero
		{"2", NumF64Ptr(nil), -10, false},           // not required
		{"3", NumF64Ptr(nil).Required(), -10, true}, // required
		{"4", NumF64Ptr(ptrF64(1)), 1, false},       // smaller than 1
		{"5", NumF64Ptr(ptrF64(1)), 2, false},
		{"6", NumF64Ptr(ptrF64(10)), 1, true},
	})
	testNumberPtrLTE(t, "NumI32Ptr", []testNumberValue{
		{"1", NumI32Ptr(ptrI32(0)), 0, false},       // test zero
		{"2", NumI32Ptr(nil), -10, false},           // not required
		{"3", NumI32Ptr(nil).Required(), -10, true}, // required
		{"4", NumI32Ptr(ptrI32(1)), 1, false},       // smaller than 1
		{"5", NumI32Ptr(ptrI32(1)), 2, false},
		{"6", NumI32Ptr(ptrI32(10)), 1, true},
	})
	testNumberPtrLTE(t, "NumI64Ptr", []testNumberValue{
		{"1", NumI64Ptr(ptrI64(0)), 0, false},       // test zero
		{"2", NumI64Ptr(nil), -10, false},           // not required
		{"3", NumI64Ptr(nil).Required(), -10, true}, // required
		{"4", NumI64Ptr(ptrI64(1)), 1, false},       // smaller than 1
		{"5", NumI64Ptr(ptrI64(1)), 2, false},
		{"6", NumI64Ptr(ptrI64(10)), 1, true},
	})
}

func testNumberPtrLTE(t *testing.T, detail string, tests []testNumberValue) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LTE(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.LTE() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}

func Test_NumberPtr_Range(t *testing.T) {
	testNumberPtrRange(t, "NumberPtr", []testNumberMinMax{
		{"1", NumberPtr(ptrNum(0)), 0, 0, false},     // test zero
		{"2", NumberPtr(nil), 1, 2, false},           // not required
		{"3", NumberPtr(nil).Required(), 1, 2, true}, // required
		{"4", NumberPtr(ptrNum(1)), 2, 4, true},
		{"5", NumberPtr(ptrNum(2)), 2, 4, false},
		{"6", NumberPtr(ptrNum(3)), 2, 4, false},
		{"7", NumberPtr(ptrNum(4)), 2, 4, false},
		{"7", NumberPtr(ptrNum(5)), 2, 4, true},
	})
	testNumberPtrRange(t, "NumF32Ptr", []testNumberMinMax{
		{"1", NumF32Ptr(ptrF32(0)), 0, 0, false},     // test zero
		{"2", NumF32Ptr(nil), 1, 2, false},           // not required
		{"3", NumF32Ptr(nil).Required(), 1, 2, true}, // required
		{"4", NumF32Ptr(ptrF32(1)), 2, 4, true},
		{"5", NumF32Ptr(ptrF32(2)), 2, 4, false},
		{"6", NumF32Ptr(ptrF32(3)), 2, 4, false},
		{"7", NumF32Ptr(ptrF32(4)), 2, 4, false},
		{"7", NumF32Ptr(ptrF32(5)), 2, 4, true},
	})
	testNumberPtrRange(t, "NumF64Ptr", []testNumberMinMax{
		{"1", NumF64Ptr(ptrF64(0)), 0, 0, false},     // test zero
		{"2", NumF64Ptr(nil), 1, 2, false},           // not required
		{"3", NumF64Ptr(nil).Required(), 1, 2, true}, // required
		{"4", NumF64Ptr(ptrF64(1)), 2, 4, true},
		{"5", NumF64Ptr(ptrF64(2)), 2, 4, false},
		{"6", NumF64Ptr(ptrF64(3)), 2, 4, false},
		{"7", NumF64Ptr(ptrF64(4)), 2, 4, false},
		{"7", NumF64Ptr(ptrF64(5)), 2, 4, true},
	})
	testNumberPtrRange(t, "NumI32Ptr", []testNumberMinMax{
		{"1", NumI32Ptr(ptrI32(0)), 0, 0, false},     // test zero
		{"2", NumI32Ptr(nil), 1, 2, false},           // not required
		{"3", NumI32Ptr(nil).Required(), 1, 2, true}, // required
		{"4", NumI32Ptr(ptrI32(1)), 2, 4, true},
		{"5", NumI32Ptr(ptrI32(2)), 2, 4, false},
		{"6", NumI32Ptr(ptrI32(3)), 2, 4, false},
		{"7", NumI32Ptr(ptrI32(4)), 2, 4, false},
		{"7", NumI32Ptr(ptrI32(5)), 2, 4, true},
	})
	testNumberPtrRange(t, "NumI64Ptr", []testNumberMinMax{
		{"1", NumI64Ptr(ptrI64(0)), 0, 0, false},     // test zero
		{"2", NumI64Ptr(nil), 1, 2, false},           // not required
		{"3", NumI64Ptr(nil).Required(), 1, 2, true}, // required
		{"4", NumI64Ptr(ptrI64(1)), 2, 4, true},
		{"5", NumI64Ptr(ptrI64(2)), 2, 4, false},
		{"6", NumI64Ptr(ptrI64(3)), 2, 4, false},
		{"7", NumI64Ptr(ptrI64(4)), 2, 4, false},
		{"7", NumI64Ptr(ptrI64(5)), 2, 4, true},
	})
}

func testNumberPtrRange(t *testing.T, detail string, tests []testNumberMinMax) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Range(tt.min, tt.max).Error(); (err != nil) != tt.wantErr {
				t.Errorf("%v.Range() error = %v, wantErr %v", detail, err, tt.wantErr)
			}
		})
	}
}
