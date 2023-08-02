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

func BenchmarkArray(b *testing.B) {
	const (
		vmin, vmax = 0, 255
		side       = 10
	)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	randInt := func(a, b int) (rv int) {
		return a + rng.Intn(b-a)
	}

	randPoint := func(n int) (x, y int) {
		return randInt(1, n), randInt(1, n)
	}

	a := array2d.New[int](side, side)
	x, y := randPoint(side)
	v := randInt(vmin, vmax)

	b.ResetTimer()

	b.Run("Set", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a.Set(x, y, v)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = a.Get(x, y)
		}
	})

	b.Run("Iter", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a.Iter(func(_, _, _ int) (next bool) {
				return true
			})
		}
	})

	b.Run("Fill", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a.Fill(func() int {
				return 1
			})
		}
	})
}
