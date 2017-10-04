package main

import (
	"fmt"

	"github.com/rodcorsi/vld"
)

type myStr string

type product struct {
	ID     string
	OldID  *string
	Descr  myStr
	Qty    int64
	Values []string
}

func main() {

	oldID := "123"
	p := product{
		ID:    "12",
		OldID: &oldID,
		Descr: "ff",
		Qty:   1,
	}

	validate := vld.New()

	_ = validate.Ok("id", vld.String(p.ID).Required().Length(1, 20).Error()) &&
		validate.Ok("old_id", vld.StrPtr(p.OldID).Length(2, 20).Error()) &&
		validate.Ok("descr", vld.String(string(p.Descr)).Length(2, 20).Error()) &&
		validate.Ok("qty", vld.NumI64(p.Qty).Range(2, 20).Error())

	fmt.Println(validate.Error())
}
