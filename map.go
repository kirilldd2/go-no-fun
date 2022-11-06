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
