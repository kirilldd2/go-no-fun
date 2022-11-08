package fun

// Map takes function that takes a value of type IN and returns a value of type OUT and a slice of type IN.
// Returns slice of type OUT.
// Example:
//
//	Map(func(n int) float64 { return float64(n) }, []int{1, 2, 3, 4})
//	// returns []float64{1., 2., 3., 4.}
func Map[IN, OUT any](fn func(IN) OUT, data []IN) []OUT {
	res := make([]OUT, len(data))

	for i, item := range data {
		res[i] = fn(item)
	}

	return res
}

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
