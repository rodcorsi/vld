package vld

import (
	"testing"
)

func Test_Number_Required(t *testing.T) {
	tests := []struct {
		name    string
		s       *numberVld
		wantErr bool
	}{
		{"1", Number(0), true},
		{"2", Number(1), false},
		{"3", NumF32(0), true},
		{"4", NumF32(1), false},
		{"5", NumF64(0), true},
		{"6", NumF64(1), false},
		{"7", NumI32(0), true},
		{"8", NumI64(1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Required().Error(); (err != nil) != tt.wantErr {
				t.Errorf("Number.Required() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Number_GT(t *testing.T) {
	tests := []struct {
		name    string
		s       *numberVld
		value   float64
		wantErr bool
	}{
		{"1", Number(0), 0, false},            // test zero
		{"2", Number(0), 10, false},           // not required
		{"3", Number(0).Required(), 10, true}, // required
		{"4", Number(1), 1, true},             // greater than 1
		{"5", Number(2), 1, false},
		{"6", Number(1), 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GT(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("Number.GT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Number_GTE(t *testing.T) {
	tests := []struct {
		name    string
		s       *numberVld
		value   float64
		wantErr bool
	}{
		{"1", Number(0), 0, false},            // test zero
		{"2", Number(0), 10, false},           // not required
		{"3", Number(0).Required(), 10, true}, // required
		{"4", Number(1), 10, true},            // greater than 1
		{"5", Number(1), 1, false},
		{"6", Number(1), 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GTE(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("Number.GTE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Number_LT(t *testing.T) {
	tests := []struct {
		name    string
		s       *numberVld
		value   float64
		wantErr bool
	}{
		{"1", Number(0), 0, false},             // test zero
		{"2", Number(0), -10, false},           // not required
		{"3", Number(0).Required(), -10, true}, // required
		{"4", Number(1), 1, true},              // smaller than 1
		{"5", Number(1), 2, false},
		{"6", Number(10), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LT(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("Number.LT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Number_LTE(t *testing.T) {
	tests := []struct {
		name    string
		s       *numberVld
		value   float64
		wantErr bool
	}{
		{"1", Number(0), 0, false},             // test zero
		{"2", Number(0), -10, false},           // not required
		{"3", Number(0).Required(), -10, true}, // required
		{"4", Number(1), 1, false},             // smaller than 1
		{"5", Number(1), 2, false},
		{"6", Number(10), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.LTE(tt.value).Error(); (err != nil) != tt.wantErr {
				t.Errorf("Number.LTE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_Number_Range(t *testing.T) {
	tests := []struct {
		name    string
		s       *numberVld
		min     float64
		max     float64
		wantErr bool
	}{
		{"1", Number(0), 0, 0, false},           // test zero
		{"2", Number(0), 1, 2, false},           // not required
		{"3", Number(0).Required(), 1, 2, true}, // required
		{"4", Number(1), 2, 4, true},
		{"5", Number(2), 2, 4, false},
		{"6", Number(3), 2, 4, false},
		{"7", Number(4), 2, 4, false},
		{"7", Number(5), 2, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Range(tt.min, tt.max).Error(); (err != nil) != tt.wantErr {
				t.Errorf("Number.Range() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
