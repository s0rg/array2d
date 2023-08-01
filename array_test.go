package array2d_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/s0rg/array2d"
)

func TestArrayBase(t *testing.T) {
	t.Parallel()

	a := array2d.New[int](15, 20)

	if w, h := a.Bounds(); w != 15 || h != 20 {
		t.Fail()
	}

	if a.Len() != 15*20 {
		t.Fail()
	}

	var (
		v  int
		ok bool
	)

	if v, ok = a.Get(0, 0); !ok {
		t.Fail()
	}

	if v != 0 {
		t.Fail()
	}

	if !a.Set(0, 0, 1) {
		t.Fail()
	}

	if v, ok = a.Get(0, 0); !ok {
		t.Fail()
	}

	if v != 1 {
		t.Fail()
	}

	if _, ok = a.Get(20, 20); ok {
		t.Fail()
	}

	if a.Set(20, 20, 10) {
		t.Fail()
	}
}

func TestArrayIterFill(t *testing.T) {
	t.Parallel()

	a := array2d.New[int](10, 10)

	a.Fill(func() int {
		return 1
	})

	var sum int

	a.Iter(func(_, _, val int) bool {
		sum += val

		return true
	})

	if sum != 100 {
		t.Fail()
	}
}

func TestArrayIterBreak(t *testing.T) {
	t.Parallel()

	a := array2d.New[int](10, 10)

	a.Fill(func() int {
		return 1
	})

	var sum int

	a.Iter(func(_, _, val int) bool {
		sum += val

		return sum <= 10
	})

	if sum != 11 {
		t.Fail()
	}
}

var benchmarkRand = rand.NewSource(time.Now().UnixNano())

func BenchmarkSet(b *testing.B) {
	const (
		vmin, vmax = 0, 255
		side       = 10
	)

	a := array2d.New[int](side, side)
	r := rand.New(benchmarkRand)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		x, y := randPoint(r, side)
		v := randInt(r, vmin, vmax)
		a.Set(x, y, v)
	}
}

func BenchmarkGet(b *testing.B) {
	const (
		side = 10
	)

	a := array2d.New[int](side, side)
	r := rand.New(benchmarkRand)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		x, y := randPoint(r, side)
		_, _ = a.Get(x, y)
	}
}

func BenchmarkIter(b *testing.B) {
	const (
		side = 10
	)

	a := array2d.New[int](side, side)
	f := func(_, _, _ int) (next bool) {
		return true
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		a.Iter(f)
	}
}

func BenchmarkFill(b *testing.B) {
	const (
		side = 10
	)

	a := array2d.New[int](side, side)
	f := func() int {
		return 1
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		a.Fill(f)
	}
}

func randPoint(r *rand.Rand, n int) (x, y int) {
	return randInt(r, 1, n), randInt(r, 1, n)
}

func randInt(r *rand.Rand, a, b int) (rv int) {
	return a + r.Intn(b-a)
}
