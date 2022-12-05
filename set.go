package fun

// Set is a container type that defines a bunch of useful methods to work with it as a set of elements.
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet creates a new set from given elements.
func NewSet[T comparable](slice []T) *Set[T] {
	data := make(map[T]struct{})
	for _, item := range slice {
		data[item] = struct{}{}
	}

	return &Set[T]{data}
}

// Slice returns a slice of current elements.
func (s *Set[T]) Slice() []T {
	slice := make([]T, s.Len())
	i := 0

	for item := range s.data {
		slice[i] = item
		i++
	}

	return slice
}

// Has indicates whether an item is present in the set.
func (s *Set[T]) Has(item T) bool {
	_, ok := s.data[item]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.data)
}

// Copy returns shallow copy of a current Set.
func (s *Set[T]) Copy() *Set[T] {
	data := make(map[T]struct{}, s.Len())
	for item := range s.data {
		data[item] = struct{}{}
	}

	return &Set[T]{data}
}

// Union returns a new Set that contains a union of callee Set and argument Set.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := make(map[T]struct{})

	for item := range s.data {
		result[item] = struct{}{}
	}

	for item := range other.data {
		result[item] = struct{}{}
	}

	return &Set[T]{result}
}

// Intersection returns a new Set that contains an intersection of callee Set and argument Set.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := make(map[T]struct{})

	for item := range s.data {
		if other.Has(item) {
			result[item] = struct{}{}
		}
	}

	return &Set[T]{result}
}

// Difference returns a new Set that contains a difference of callee Set and argument Set.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := make(map[T]struct{})

	for item := range s.data {
		if !other.Has(item) {
			result[item] = struct{}{}
		}
	}

	return &Set[T]{result}
}

// Disjoint returns true if the callee Set has no common elements with the other Set.
func (s *Set[T]) Disjoint(other *Set[T]) bool {
	if s.Intersection(other).Len() == 0 {
		return true
	} else {
		return false
	}
}

// Subset returns true if the callee Set is a subset of the other Set.
func (s *Set[T]) Subset(other *Set[T]) bool {
	for item := range s.data {
		if !other.Has(item) {
			return false
		}
	}

	return true
}

// Superset returns true if the callee Set is a superset of the other Set.
func (s *Set[T]) Superset(other *Set[T]) bool {
	return other.Subset(s)
}

// Equals returns true if the two sets are equal.
func (s *Set[T]) Equals(other *Set[T]) bool {
	return Equal(s.Slice(), other.Slice())
}

// SymmetricDifference returns a new Set with elements in either the callee Set or other but not both.
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	result := make(map[T]struct{})

	for item := range s.data {
		if !other.Has(item) {
			result[item] = struct{}{}
		}
	}

	for item := range other.data {
		if !s.Has(item) {
			result[item] = struct{}{}
		}
	}

	return &Set[T]{result}
}

// Add adds items to the set. Returns itself.
func (s *Set[T]) Add(items ...T) *Set[T] {
	for _, item := range items {
		s.data[item] = struct{}{}
	}

	return s
}

// Update adds items to the set from given sets. Returns itself.
func (s *Set[T]) Update(others ...*Set[T]) *Set[T] {
	for _, other := range others {
		for item := range other.data {
			s.data[item] = struct{}{}
		}
	}

	return s
}

// Remove removes items from the set. Returns true if item was in the set.
func (s *Set[T]) Remove(item T) bool {
	ok := s.Has(item)

	delete(s.data, item)

	return ok
}
