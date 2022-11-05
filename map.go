package fun

// Map takes function that takes a value of type IN and returns a value of type OUT
// and a slice of type IN.
// Returns slice of type out
// Example: Map(func(n int) float)
func Map[IN, OUT any](fn func(IN) OUT, data []IN) []OUT {
	res := make([]OUT, len(data))

	for i := range data {
		res[i] = fn(data[i])
	}

	return res
}

func MapErr[IN, OUT any](fn func(IN) (OUT, error), data []IN) ([]OUT, error) {
	res := make([]OUT, len(data))

	for i := range data {
		mapped, err := fn(data[i])
		if err != nil {
			return nil, err
		}
		res[i] = mapped
	}

	return res, nil
}
