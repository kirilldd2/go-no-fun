package fast

import (
	"sync"
)

// Map takes function that takes a value of type IN and returns a value of type OUT and a slice of type IN.
// Returns slice of type OUT.
func Map[IN, OUT any](fn func(IN) OUT, data []IN, workers uint8) []OUT {
	length := len(data)
	if length == 0 {
		return []OUT{}
	}

	batch := length / int(workers)
	res := make([]OUT, length)

	var wg sync.WaitGroup
	for i := uint8(0); i < workers; i++ {
		end := (int(i) + 1) * batch
		if i == workers-1 {
			end = length
		}

		wg.Add(1)
		go func(data []IN, start int) {
			for i, item := range data {
				res[i+start] = fn(item)
			}

			wg.Done()
		}(data[int(i)*batch:end], int(i)*batch)
	}

	wg.Wait()

	return res
}
