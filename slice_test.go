package fun_test

import (
	"reflect"
	"testing"

	fun "github.com/kirilldd2/go-no-fun"
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

func TestReversed(t *testing.T) {
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
			if res := fun.Reversed(test.inp); !fun.Equal(res, test.out) {
				t.Errorf("got %v, wanted %v", res, test.out)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	var tests = []struct {
		name      string
		inp       []int
		find, res int
	}{
		{"success", []int{1, 2, 3, 4}, 2, 1},
		{"not found", []int{1, 2, 3, 4}, 10, -1},
		{"last", []int{1, 2, 3, 4, 100}, 100, 4},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if res := fun.Index(test.inp, test.find); res != test.res {
				t.Errorf("got %v, wanted %v", res, test.res)
			}
		})
	}
}

func TestIndexAB(t *testing.T) {
	var tests = []struct {
		name            string
		inp             []int
		a, b, find, res int
		err             bool
	}{
		{"success full", []int{1, 2, 3, 4}, 0, 4, 2, 1, false},
		{"success", []int{1, 2, 3, 4}, 0, 3, 2, 1, false},
		{"not found", []int{1, 2, 3, 4}, 0, 2, 3, -1, false},
		{"not found before", []int{1, 2, 3, 4}, 1, 3, 1, -1, false},
		{"error a==b", []int{1, 2, 3, 4}, 1, 1, 1, -1, true},
		{"error a<0", []int{1, 2, 3, 4}, -1, 1, 1, -1, true},
		{"error b>len(slice)", []int{1, 2, 3, 4}, 1, 5, 1, -1, true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res, err := fun.IndexAB(test.inp, test.find, test.a, test.b)
			if err != nil && test.err == false {
				t.Errorf("unexpected error on test %+v", test)
			} else if test.err == false && res != test.res {
				t.Errorf("got %v, wanted %v", res, test.res)
			}
		})
	}
}
