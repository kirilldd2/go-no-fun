package fun

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](slice []T) *Set[T] {
	data := make(map[T]struct{})
	for _, item := range slice {
		data[item] = struct{}{}
	}

	return &Set[T]{data}
}

func (s *Set[T]) Hash() map[T]struct{} {
	return s.data
}

func (s *Set[T]) Slice() []T {
	slice := make([]T, len(s.data))
	i := 0

	for item := range s.data {
		slice[i] = item
		i++
	}

	return slice
}

func (s *Set[T]) Has(item T) bool {
	_, ok := s.Hash()[item]
	return ok
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	result := make(map[T]struct{})

	for item := range s.Hash() {
		result[item] = struct{}{}
	}

	for item := range s2.Hash() {
		result[item] = struct{}{}
	}

	return &Set[T]{result}
}

func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	result := make(map[T]struct{})

	for item := range s.Hash() {
		if _, ok := s2.Hash()[item]; ok {
			result[item] = struct{}{}
		}
	}

	return &Set[T]{result}
}

func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	result := make(map[T]struct{})

	for item := range s.Hash() {
		if _, ok := s2.Hash()[item]; !ok {
			result[item] = struct{}{}
		}
	}

	return &Set[T]{result}
}
