package array2d

import "testing"

func TestArrayBase(t *testing.T) {
	t.Parallel()

	a := New[int](15, 20)

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

	a := New[int](10, 10)

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

	a := New[int](10, 10)

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
