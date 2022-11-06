package fun

// Reduce applies function fn of two args cumulatively to the items of iterable, from left to right,
// to reduce the iterable to a single value, starting from init value.
//
// Example:
//
//	// get 1**2 + 2**2 + 3**2 + 4**2
//	Reduce(func(acc float64, item int) { return acc + math.Pow(float64(item), 2) }, []int(1, 2, 3, 4), 1)
func Reduce[A, I any](fn func(acc A, item I) A, iterable []I, init A) A {
	for i := range iterable {
		init = fn(init, iterable[i])
	}

	return init
}
