package fun

import "errors"

var errOutOfBoundaries = errors.New("incorrect boundaries of slice")

func Reverse[T any](slice []T) {
	l := len(slice)
	for i := 0; i < l/2; i++ {
		slice[i], slice[l-i-1] = slice[l-i-1], slice[i]
	}
}

func Reversed[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	Reverse(newSlice)

	return newSlice
}

func Index[T comparable](slice []T, value T) int {
	for i, item := range slice {
		if item == value {
			return i
		}
	}

	return -1
}

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
