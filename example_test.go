package fun_test

import (
	"fmt"
	"math"

	fun "github.com/kirilldd2/go-no-fun"
)

func ExampleMap() {
	result := fun.Map(func(n int) int64 { return int64(n * n) }, []int{1, 2, 3, 4})

	fmt.Println(result)

	// Output: [1 4 9 16]
}

func ExampleReduce() {
	result := fun.Reduce(
		func(acc float64, item int) float64 { return acc + math.Pow(float64(item), 2) },
		[]int{1, 2, 3, 4}, 0,
	)

	fmt.Println(result)

	// Output: 30
}

func ExampleFilter() {
	result := fun.Filter(func(item string) bool { return item[0] == 'c' }, []string{"cunt", "sucka", "car", "tree"})

	fmt.Println(result)

	// Output: [cunt car]
}

func ExampleAll() {
	result1 := fun.All([]string{"hello", "gentlemen", "!"})
	result2 := fun.All([]string{"hello", "gentlemen", ""})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleAny() {
	result1 := fun.Any([]string{"", "a", ""})
	result2 := fun.Any([]string{"", "", ""})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleMin() {
	type Test struct {
		A int
	}

	result1 := fun.Min(fun.Less[int], []int{1, 2, -1, 3}...)
	result2 := fun.Min(
		func(less, bigger Test) bool { return less.A < bigger.A },
		[]Test{{1}, {3}, {-2}, {5}}...,
	)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// -1
	// {-2}
}

func ExampleMax() {
	type Test struct {
		A int
	}

	result1 := fun.Max(fun.Less[int], []int{1, 2, -1, 3}...)
	result2 := fun.Max(
		func(less, bigger Test) bool { return less.A < bigger.A },
		[]Test{{1}, {3}, {-2}, {5}}...,
	)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 3
	// {5}
}

func ExampleEqual() {
	result1 := fun.Equal([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})
	result2 := fun.Equal([]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5})
	result3 := fun.Equal([]int{1, 2, 3, 4}, []int{1, 2, 3, 1})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleReverse() {
	slice := []string{"!", "gentlemen", "Hello"}

	fun.Reverse(slice)

	fmt.Println(slice)

	// Output: [Hello gentlemen !]
}

func ExampleReversed() {
	slice := []string{"!", "gentlemen", "Hello"}
	newSlice := fun.Reversed(slice)

	fmt.Println(slice)
	fmt.Println(newSlice)

	// Output:
	// [! gentlemen Hello]
	// [Hello gentlemen !]
}

func ExampleIndex() {
	// Note: only first encounter is returned
	result := fun.Index([]string{"Hello", "gentlemen", "!", "gentlemen"}, "gentlemen")

	fmt.Println(result)

	// Output: 1
}

func ExampleIndexAB() {
	// Note: only first encounter is returned
	result, _ := fun.IndexAB([]string{"Hello", "gentlemen", "!", "gentlemen"}, "gentlemen", 2, 4)

	fmt.Println(result)

	// Output: 3
}

func ExampleSum() {
	resultInts := fun.Sum([]int{6, -5, 10})
	resultStr := fun.Sum([]string{"Hello,", " gentlemen", "!"})

	fmt.Println(resultInts)
	fmt.Println(resultStr)

	// Output:
	// 11
	// Hello, gentlemen!
}
