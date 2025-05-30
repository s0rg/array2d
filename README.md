[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/array2d)](https://pkg.go.dev/github.com/s0rg/array2d)
[![License](https://img.shields.io/github/license/s0rg/array2d)](https://github.com/s0rg/array2d/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/array2d)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/array2d?sort=semver)](https://github.com/s0rg/array2d/tags)

[![CI](https://github.com/s0rg/array2d/workflows/ci/badge.svg)](https://github.com/s0rg/array2d/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/array2d)](https://goreportcard.com/report/github.com/s0rg/array2d)
[![Maintainability](https://qlty.sh/badges/07ef5116-fc80-4eea-9c04-769a32323112/maintainability.svg)](https://qlty.sh/gh/s0rg/projects/array2d)
[![Code Coverage](https://qlty.sh/badges/07ef5116-fc80-4eea-9c04-769a32323112/test_coverage.svg)](https://qlty.sh/gh/s0rg/projects/array2d)
![Issues](https://img.shields.io/github/issues/s0rg/array2d)

# array2d

Generic 2D array


# features

- zero-alloc
- 100% test coverage


# example

```go
import (
    "fmt"

    "github.com/s0rg/array2d"
)

func main() {
    // create new 2D array with given type and dimensions (width x height)
    a := array2d.New[int](15, 20)

    // fill with ones
    a.Fill(func() (value int) {
        return 1
    })

    // set value at given coords
    a.Set(1, 2, 100)

    // get values
    t1, ok := a.Get(1, 1)
    if !ok {
        panic("nothing at 1x1")
    }

    t2, ok := a.Get(1, 2)
    if !ok {
        panic("nothing at 1x2")
    }

    fmt.Println(t1, t2) // should print: 1, 100
}
```


# benchmarks

```
goos: linux
goarch: amd64
pkg: github.com/s0rg/array2d
cpu: AMD Ryzen 5 5500U with Radeon Graphics
BenchmarkArray/Set-12           1000000000         1.002 ns/op        0 B/op          0 allocs/op
BenchmarkArray/Get-12           927178446          1.250 ns/op        0 B/op          0 allocs/op
BenchmarkArray/Iter-12          39491674           30.68 ns/op        0 B/op          0 allocs/op
BenchmarkArray/Fill-12          20840200           56.84 ns/op        0 B/op          0 allocs/op
PASS
ok      github.com/s0rg/array2d 4.897s
```
