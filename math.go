package fun

import "golang.org/x/exp/constraints"

// Less is default comparator function for builtin data-types for Min or Max.
func Less[T constraints.Ordered](less, big T) bool {
	return less < big
}

// Min returns minimum out of values.
// To compare values "less" function needs to be provided. Basic example is Less.
// If no values provided returns zero value of T.
func Min[T any](less func(less, big T) bool, values ...T) T {
	var min T

	if len(values) == 0 {
		return min
	}

	min = values[0]
	for i := 1; i < len(values); i++ {
		if less(values[i], min) {
			min = values[i]
		}
	}

	return min
}

// Max returns maximum out of values.
// To compare values "less" function needs to be provided. Basic example is Less.
// If no values provided returns zero value of T.
func Max[T any](less func(less, big T) bool, values ...T) T {
	var max T

	if len(values) == 0 {
		return max
	}

	max = values[0]
	for i := 1; i < len(values); i++ {
		if less(max, values[i]) {
			max = values[i]
		}
	}

	return max
}
