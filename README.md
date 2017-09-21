# vld

[![Build Status](https://travis-ci.org/rodcorsi/vld.svg?branch=master)](https://travis-ci.org/rodcorsi/vld)
[![Coverage Status](https://coveralls.io/repos/github/rodcorsi/vld/badge.svg?branch=master)](https://coveralls.io/github/rodcorsi/vld?branch=master)

Package written in Go for validation of string, numbers and structs without use of tags

**Vld may not be considered stable just yet!**

## Installation

You'll need Go installed

```bash
go get github.com.br/rodcorsi/vld
```

## Usage

```go
oldID := "123"
p := Product{
    ID:    "12",
    OldID: &oldID,
    Descr: "ff",
    Qty:   10,
}

validate := vld.New()

// use && to stop verification if has one error
// you can use || to check all errors
_ = validate.Ok("id", vld.String(p.ID).Required().Length(1, 20).Error()) &&
    validate.Ok("old_id", vld.StrPtr(p.OldID).Length(2, 20).Error()) &&
    validate.Ok("descr", vld.String(string(p.Descr)).Length(2, 20).Error()) &&
    validate.Ok("qty", vld.NumI64(p.Qty).Range(2, 20).Error())

if err := validate.Error(); err != nil {
    fmt.Println(err)
}
```

## Benchmark

Benchmark comparing other validator

```bash
$ go test -run=NONE -bench=. -benchmem ./example/
goos: linux
goarch: amd64
pkg: github.com/rodcorsi/vld/example
BenchmarkGovalidator-4            100000             17955 ns/op            2424 B/op         39 allocs/op
BenchmarkValidator-4             1000000              1045 ns/op              64 B/op          2 allocs/op
BenchmarkOzzo-4                   200000              7863 ns/op            2104 B/op         50 allocs/op
BenchmarkFourSigma-4              500000              2103 ns/op             688 B/op         25 allocs/op
BenchmarkVld-4                  100000000               16.7 ns/op             0 B/op          0 allocs/op
```