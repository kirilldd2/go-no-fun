/*
Package fun provides simple and useful functional (and other) utilities that often has to be made by hand.
Utilizes power of Go 1.18 generics.
*/
package fun

// Map takes function that takes a value of type IN and returns a value of type OUT and a slice of type IN.
// Returns slice of type OUT.
func Map[IN, OUT any](fn func(IN) OUT, data []IN) []OUT {
	res := make([]OUT, len(data))

	for i, item := range data {
		res[i] = fn(item)
	}

	return res
}

// Reduce applies function fn of two args cumulatively to the items of iterable, from left to right,
// to reduce the iterable to a single value, starting from init value.
func Reduce[A, I any](fn func(acc A, item I) A, iterable []I, init A) A {
	for i := range iterable {
		init = fn(init, iterable[i])
	}

	return init
}

// Filter returns new slice of items, that fn(item) == true.
func Filter[T any](fn func(item T) bool, iterable []T) []T {
	filtered := make([]T, 0, len(iterable))

	for _, item := range iterable {
		if fn(item) {
			filtered = append(filtered, item)
		}
	}

	// reallocate to free excess memory
	adjusted := make([]T, len(filtered))
	copy(adjusted, filtered)

	return adjusted
}

// Any returns true if there's any non-nil values in iterable, false otherwise.
func Any[T comparable](iterable []T) bool {
	var zero T
	for _, item := range iterable {
		if item != zero {
			return true
		}
	}

	return false
}

// All returns true if all the values in iterable are non-nil, false otherwise.
func All[T comparable](iterable []T) bool {
	var zero T
	for _, item := range iterable {
		if item == zero {
			return false
		}
	}

	return true
}
