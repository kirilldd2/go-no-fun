package fun_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	. "github.com/Ioloman/go-no-fun"
)

func IntToFloat64(n int) float64 { return math.Sqrt(float64(n)) }

func createRandomSlice(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	slice := make([]int, n)
	for i := range slice {
		slice[i] = r.Intn(math.MaxInt)
	}

	return slice
}

func TestMap(t *testing.T) {

	t.Run("ints to float64 by sqrt", func(t *testing.T) {
		inp := createRandomSlice(1000)
		want := make([]float64, 1000)
		for i := range inp {
			want[i] = IntToFloat64(inp[i])
		}
		result := Map(IntToFloat64, inp)
		if len(result) != len(inp) {
			t.Errorf("result len = %d, input's len = %d", len(result), len(inp))
		}
		for i := range result {
			if result[i] != want[i] {
				t.Errorf("on index = %d got %f, but wanted %f", i, result[i], want[i])
			}
		}
	})

	t.Run("empty input slice", func(t *testing.T) {
		var inp []int
		result := Map(IntToFloat64, inp)
		if len(result) != 0 {
			t.Error("result for input empty slice is not empty slice")
		}
	})
}

func BenchmarkMap(b *testing.B) {
	inp := createRandomSlice(1000)

	b.Run("Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Map(IntToFloat64, inp)
		}
	})

	b.Run("For loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out := make([]float64, len(inp))
			for j := range inp {
				out[j] = IntToFloat64(inp[j])
			}
		}
	})
}
