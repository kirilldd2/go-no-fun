package fun

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func IntToFloat64(n int) float64 { return math.Sqrt(float64(n)) }

func TestMap(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	t.Run("ints to float64 by sqrt", func(t *testing.T) {
		inp := make([]int, 1000)
		want := make([]float64, 1000)
		for i := range inp {
			inp[i] = r.Intn(math.MaxInt)
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
}

func BenchmarkMap(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := 0; i < b.N; i++ {
		inp := make([]int, b.N)
		for i := range inp {
			inp[i] = r.Intn(math.MaxInt)
		}
		Map(IntToFloat64, inp)
	}
}

func BenchmarkLoop(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for i := 0; i < b.N; i++ {
		inp := make([]int, b.N)
		for i := range inp {
			inp[i] = r.Intn(math.MaxInt)
		}
		out := make([]float64, len(inp))
		for j := range inp {
			out[j] = IntToFloat64(inp[j])
		}
	}
}
