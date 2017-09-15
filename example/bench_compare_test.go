package main

import (
	"errors"
	"testing"

	"github.com/FourSigma/validate"
	"github.com/FourSigma/validate/str"
	"github.com/asaskevich/govalidator"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/rodcorsi/vld"
	validator "gopkg.in/go-playground/validator.v9"
)

type product struct {
	ID    string  `valid:"required,length(1|20)"`
	OldID *string `valid:"length(2|20)"`
	Descr myStr   `valid:"length(2|20)"`
	Qty   int64   `valid:"range(1|100)"`
}

type productv9 struct {
	ID    string  `validate:"required,gte=1,lte=20"`
	OldID *string `validate:"gte=2,lte=20"`
	Descr myStr   `validate:"gte=2,lte=20"`
	Qty   int64   `validate:"gte=1,lte=100"`
}

func BenchmarkGovalidator(b *testing.B) {
	oldID := "123"
	p := product{
		ID:    "12",
		OldID: &oldID,
		Descr: "ff",
		Qty:   10,
	}
	for n := 0; n < b.N; n++ {
		govalidator.ValidateStruct(p)
	}
}

func BenchmarkValidator(b *testing.B) {
	oldID := "123"
	p := productv9{
		ID:    "12",
		OldID: &oldID,
		Descr: "ff",
		Qty:   10,
	}
	validate := validator.New()
	for n := 0; n < b.N; n++ {
		validate.Struct(p)
	}
}

func BenchmarkOzzo(b *testing.B) {
	oldID := "123"
	p := product{
		ID:    "12",
		OldID: &oldID,
		Descr: "ff",
		Qty:   10,
	}

	for n := 0; n < b.N; n++ {
		validation.ValidateStruct(&p,
			validation.Field(&p.ID, validation.Required, validation.Length(1, 20)),
			validation.Field(&p.OldID, validation.Length(2, 20)),
			validation.Field(&p.Descr, validation.Length(2, 20)),
			validation.Field(&p.Qty, validation.Min(1), validation.Max(100)),
		)
	}
}

func BenchmarkFourSigma(b *testing.B) {
	oldID := "123"
	p := product{
		ID:    "12",
		OldID: &oldID,
		Descr: "ff",
		Qty:   10,
	}
	for n := 0; n < b.N; n++ {
		desc := string(p.Descr)
		validate.Check(
			validate.String(&p.ID).Validate(str.MinLen(1), str.MaxLen(20)).Required(),
			validate.String(p.OldID).Validate(str.MinLen(2), str.MaxLen(20)),
			validate.String(&desc).Validate(str.MinLen(2), str.MaxLen(20)),
			validate.String(&desc).Validate(str.MinLen(2), str.MaxLen(20)),
		)
	}
}

func BenchmarkVld(b *testing.B) {

	oldID := "123"
	p := product{
		ID:    "12",
		OldID: &oldID,
		Descr: "ff",
		Qty:   10,
	}

	for i := 0; i < b.N; i++ {
		validate := vld.New()
		_ = validate.Init("id").String(p.ID).Required().Length(1, 20).Ok() &&
			validate.Init("old_id").StrPtr(p.OldID).Length(2, 20).Ok() &&
			validate.Init("descr").String(string(p.Descr)).Length(2, 20).Ok() &&
			validate.Init("qty").NumI64(p.Qty).Range(2, 20).Ok()
	}
}

func BenchmarkRaw(b *testing.B) {
	oldID := "123"
	p := product{
		ID:    "12",
		OldID: &oldID,
		Descr: "ff",
		Qty:   10,
	}
	valid := func() []error {
		var errs []error
		if p.ID == "" || len(p.ID) < 1 || len(p.ID) > 20 {
			errs = append(errs, errors.New("invalid id"))
		}
		if p.OldID != nil {
			if len(*p.OldID) < 2 || len(*p.OldID) > 20 {
				errs = append(errs, errors.New("invalid old_id"))
			}
		}
		if p.Descr != "" {
			if len(p.Descr) < 2 || len(p.Descr) > 20 {
				errs = append(errs, errors.New("invalid descr"))
			}
		}
		if p.Qty != 0 {
			if p.Qty < 1 || p.Qty > 100 {
				errs = append(errs, errors.New("invalid qty"))
			}
		}
		return errs
	}

	for n := 0; n < b.N; n++ {
		valid()
	}
}
