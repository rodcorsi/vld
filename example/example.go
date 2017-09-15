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
	_ = validate.Init("id").String(p.ID).Required().Length(1, 20).Ok() &&
		validate.Init("old_id").StrPtr(p.OldID).Length(2, 20).Ok() &&
		validate.Init("descr").String(string(p.Descr)).Length(2, 20).Ok() &&
		validate.Init("qty").NumI64(p.Qty).Range(2, 20).Ok()

	if err := validate.Err(); err != nil {
		fmt.Println(err)
	}
}
