package fixtures

import (
	"math"
	"math/rand"
	"time"
)

func IntToFloat64(n int) float64 { return math.Sqrt(float64(n)) }

func CreateRandomSlice(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	slice := make([]int, n)
	for i := range slice {
		slice[i] = r.Intn(math.MaxInt)
	}

	return slice
}
