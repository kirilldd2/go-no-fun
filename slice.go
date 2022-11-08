package fun

import "errors"

var errOutOfBoundaries = errors.New("incorrect boundaries of slice")

// Equal compares two flat slices that contain comparable type.
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Reverse reverses slice in-place.
func Reverse[T any](slice []T) {
	l := len(slice)
	for i := 0; i < l/2; i++ {
		slice[i], slice[l-i-1] = slice[l-i-1], slice[i]
	}
}

// Reversed returns reversed slice, original slice is not mutated.
func Reversed[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	copy(newSlice, slice)
	Reverse(newSlice)

	return newSlice
}

// Index returns the index of first encountered "value" in slice.
// If value is not found -1 is returned.
func Index[T comparable](slice []T, value T) int {
	for i, item := range slice {
		if item == value {
			return i
		}
	}

	return -1
}

// IndexAB  returns the index of first encountered "value" in slice within interval [a, b).
// If [a, b) interval is incorrect returns 0 and error.
func IndexAB[T comparable](slice []T, value T, a, b int) (int, error) {
	if a < 0 || b <= a || b > len(slice) {
		return 0, errOutOfBoundaries
	}

	slice = slice[a:b]
	for i, item := range slice {
		if item == value {
			return i, nil
		}
	}

	return -1, nil
}
