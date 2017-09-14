package main

import (
	"fmt"

	"github.com/rodcorsi/vld"
)

type myStr string

type Product struct {
	ID    string
	OldID *string
	Descr myStr
	Qty   int64
}

func main() {
	oldID := "123"
	p := Product{
		ID:    "12",
		OldID: &oldID,
		Descr: "ff",
		Qty:   10,
	}

	validate := vld.New()
	_ = validate.String("id", p.ID).Required().Length(1, 20).Ok() &&
		validate.StrPtr("old_id", p.OldID).Length(2, 20).Ok() &&
		validate.String("descr", string(p.Descr)).Length(2, 20).Ok() &&
		validate.NumI64("qty", p.Qty).Range(2, 20).Ok()

	if err := validate.Err(); err != nil {
		fmt.Println("error: " + err.Error())
	}
}
