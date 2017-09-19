# vld

Package written in Go for validation of string, numbers and structs without use of tags

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

if validate.HasError() {
    fmt.Println(validate.Error())
}
```

## Benchmark

```bash
$ go test -run=NONE -bench=. -benchmem ./example/
goos: windows
goarch: amd64
pkg: github.com/rodcorsi/vld/example
BenchmarkGovalidator-4            100000             18983 ns/op            3072 B/op         48 allocs/op
BenchmarkValidator-4             2000000               973 ns/op              64 B/op          2 allocs/op
BenchmarkOzzo-4                   300000              5310 ns/op            2104 B/op         50 allocs/op
BenchmarkFourSigma-4             1000000              1395 ns/op             688 B/op         25 allocs/op
BenchmarkVld-4                  50000000                27.0 ns/op             0 B/op          0 allocs/op
```