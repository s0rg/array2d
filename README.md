[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/array2d)](https://pkg.go.dev/github.com/s0rg/array2d)
[![License](https://img.shields.io/github/license/s0rg/array2d)](https://github.com/s0rg/array2d/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/array2d)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/array2d?sort=semver)](https://github.com/s0rg/array2d/tags)

[![CI](https://github.com/s0rg/array2d/workflows/ci/badge.svg)](https://github.com/s0rg/array2d/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/array2d)](https://goreportcard.com/report/github.com/s0rg/array2d)
[![Maintainability](https://api.codeclimate.com/v1/badges/54e42106bc739ae75de9/maintainability)](https://codeclimate.com/github/s0rg/array2d/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/54e42106bc739ae75de9/test_coverage)](https://codeclimate.com/github/s0rg/array2d/test_coverage)
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
BenchmarkSet-12     	29237254	       39.97 ns/op	      0 B/op	      0 allocs/op
BenchmarkGet-12     	44087026	       27.37 ns/op	      0 B/op	      0 allocs/op
BenchmarkIter-12    	20285022	       55.28 ns/op	      0 B/op	      0 allocs/op
BenchmarkFill-12    	19116193	       57.65 ns/op	      0 B/op	      0 allocs/op
PASS
ok  	github.com/s0rg/array2d	5.812s
```
