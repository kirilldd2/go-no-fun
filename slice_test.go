package fun_test

import (
	"reflect"
	"testing"

	"github.com/Ioloman/go-no-fun"
)

func TestEqual(t *testing.T) {
	var tests = []struct {
		name  string
		a, b  []int
		equal bool
	}{
		{"equal", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, true},
		{"unequal_last", []int{1, 2, 3, 7}, []int{1, 2, 3, 4}, false},
		{"unequal", []int{1, 2, 3, 4, 100}, []int{100, 4, 3, 2, 1}, false},
		{"unequal_len", []int{1, 2, 3, 4, 100}, []int{1, 2, 3, 4, 100, 1}, false},
		{"empty", []int{}, []int{}, true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if fun.Equal(test.a, test.b) != test.equal {
				t.Errorf("test %+v wrong", test)
			}
		})
	}
}

func BenchmarkEqual(b *testing.B) {
	var eq []int

	n := 10000
	slice := createRandomSlice(n)

	copy(eq, slice)

	b.Run("Equal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fun.Equal(slice, eq)
		}
	})

	b.Run("reflect.DeepEqual", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = reflect.DeepEqual(slice, eq)
		}
	})
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		name     string
		inp, out []int
	}{
		{"even", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"odd", []int{1, 2, 3, 4, 100}, []int{100, 4, 3, 2, 1}},
		{"empty", []int{}, []int{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			fun.Reverse(test.inp)
			if !fun.Equal(test.inp, test.out) {
				t.Errorf("test %+v failed", test)
			}
		})
	}
}
