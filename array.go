package array2d

// Array is a generic 2D array.
type Array[T any] struct {
	v    []T
	w, h int
}

// New creates empty array with given dimensions.
func New[T any](w, h int) (a Array[T]) {
	return Array[T]{
		w: w,
		h: h,
		v: make([]T, w*h),
	}
}

// Bounds returns width and height of array.
func (a Array[T]) Bounds() (w, h int) {
	return a.w, a.h
}

// Len returns length of array.
func (a Array[T]) Len() (rv int) {
	return len(a.v)
}

// Get returns value at given coords, if any.
func (a Array[T]) Get(x, y int) (rv T, ok bool) {
	if !a.valid(x, y) {
		return
	}

	return a.v[a.index(x, y)], true
}

// Set sets value at given coords.
func (a *Array[T]) Set(x, y int, v T) (ok bool) {
	if !a.valid(x, y) {
		return
	}

	a.v[a.index(x, y)] = v

	return true
}

// Iter iterates over array items.
func (a *Array[T]) Iter(it func(x, y int, v T) bool) {
	var x, y int

	for i := 0; i < len(a.v); i++ {
		x, y = i%a.w, i/a.w

		if !it(x, y, a.v[i]) {
			return
		}
	}
}

// Fill fills array with given function.
func (a *Array[T]) Fill(fn func() T) {
	for i := 0; i < len(a.v); i++ {
		a.v[i] = fn()
	}
}

func (a *Array[T]) valid(x, y int) (ok bool) {
	if (0 > x || x >= a.w) || (0 > y || y >= a.h) {
		return
	}

	return true
}

func (a *Array[T]) index(x, y int) (rv int) {
	return x + y*a.w
}
