package fast_test

import (
	fun "github.com/kirilldd2/go-no-fun/fast"
	"github.com/kirilldd2/go-no-fun/fixtures"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("ints to float64 by sqrt", func(t *testing.T) {
		inp := fixtures.CreateRandomSlice(1000)
		want := make([]float64, 1000)
		for i := range inp {
			want[i] = fixtures.IntToFloat64(inp[i])
		}
		result := fun.Map(fixtures.IntToFloat64, inp, 2)
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
		result := fun.Map(fixtures.IntToFloat64, inp, 2)
		if len(result) != 0 {
			t.Error("result for input empty slice is not empty slice")
		}
	})
}
