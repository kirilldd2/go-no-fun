package fun_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	. "github.com/kirilldd2/go-no-fun"
)

func TestMin(t *testing.T) {
	t.Run("test on map[string]int", func(t *testing.T) {
		var m = []map[string]int{{"num": 25}, {"num": 2}, {"num": 3}}
		result := Min(func(l, b map[string]int) bool { return l["num"] < b["num"] }, m...)
		if result["num"] != 2 {
			t.Errorf("Min wrong result: %d", result["num"])
		}
	})

	t.Run("test with Less[int]", func(t *testing.T) {
		result := Min[int](Less[int], 8, 9, -6, 1000)
		if result != -6 {
			t.Errorf("Min wrong result: %d", result)
		}
	})

	t.Run("test empty", func(t *testing.T) {
		result := Min[int](Less[int])
		if result != 0 {
			t.Error("result must be 0")
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("test on map[string]int", func(t *testing.T) {
		var m = []map[string]int{{"num": 25}, {"num": 2}, {"num": 3}}
		result := Max(func(l, b map[string]int) bool { return l["num"] < b["num"] }, m...)
		if result["num"] != 25 {
			t.Errorf("Max wrong result: %d", result["num"])
		}
	})

	t.Run("test with Less[int]", func(t *testing.T) {
		result := Max[int](Less[int], 8, 9, -6, 1000)
		if result != 1000 {
			t.Errorf("Max wrong result: %d", result)
		}
	})

	t.Run("test empty", func(t *testing.T) {
		result := Max[int](Less[int])
		if result != 0 {
			t.Error("result must be 0")
		}
	})
}

func BenchmarkMin(b *testing.B) {
	random := rand.New(rand.NewSource(time.Now().UnixMilli()))

	inp := make([]int, 10000)
	for i := range inp {
		inp[i] = random.Intn(math.MaxInt)
	}

	for i := 0; i < b.N; i++ {
		Min(Less[int], inp...)
	}
}
