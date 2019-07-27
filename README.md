# gen-iter
Implements iterator for specified type with code generation.

## Installation

```bash
    go get -u github.com/BoRuDar/gen-iter
    go install github.com/BoRuDar/gen-iter
```

## Usage
- From terminal
```bash
    gen-iter -type int -pkg iter 
```
- With go generate. Add next comment into your file:
```go
    //go:generate gen-iter -t int -p iter
```

This will generate `IntIterator_gen.go` (in current directory) for `int` type and package name `iter`.