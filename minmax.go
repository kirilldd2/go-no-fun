package fun

import "golang.org/x/exp/constraints"

// Less is default comparator function for builtin data-types for Min or Max
func Less[T constraints.Ordered](less, big T) bool {
	return less < big
}

// Min returns minimum out of values.
// To compare values "less" function needs to be provided. Basic example is Less.
// If no values provided returns zero value of T.
// Example:
//
//	Min(func(l, b map[string]int) { return l["num"] < b["num"] }, []map[string]int{{"num": 2}, {"num": 3}}...)
func Min[T any](less func(less, big T) bool, values ...T) T {
	if len(values) == 0 {
		var nil T
		return nil
	}

	min := values[0]
	for i := 1; i < len(values); i++ {
		if less(values[i], min) {
			min = values[i]
		}
	}

	return min
}

// Max - refer to Min for docs
func Max[T any](less func(less, big T) bool, values ...T) T {
	if len(values) == 0 {
		var nil T
		return nil
	}

	max := values[0]
	for i := 1; i < len(values); i++ {
		if less(max, values[i]) {
			max = values[i]
		}
	}

	return max
}
