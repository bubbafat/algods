package algods

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Set[T constraints.Ordered] struct {
	data []T
}

func (s *Set[T]) insertSorted(value T) bool {
	i, found := slices.BinarySearch(s.data, value)

	if !found {
		var empty T
		s.data = append(s.data, empty)

		copy(s.data[i+1:], s.data[i:])
		s.data[i] = value
		return true
	}

	return false
}

func (s *Set[T]) Count() int {
	return len(s.data)
}

func (s *Set[T]) Contains(value T) bool {
	_, found := slices.BinarySearch(s.data, value)
	return found
}

func (s *Set[T]) Add(value T) bool {
	return s.insertSorted(value)
}

func (s *Set[T]) Remove(value T) bool {
	i, found := slices.BinarySearch(s.data, value)

	if found {
		s.data = append(s.data[:i], s.data[i+1:]...)
	}

	return found
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := new(Set[T])

	for _, x := range s.data {
		result.insertSorted(x)
	}

	for _, x := range other.data {
		result.insertSorted(x)
	}

	return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := new(Set[T])

	for _, x := range s.data {
		if other.Contains(x) {
			result.insertSorted(x)
		}
	}

	for _, x := range other.data {
		if s.Contains(x) {
			result.insertSorted(x)
		}
	}

	return result
}

func (s *Set[T]) ForEach(callback func(value T) bool) {
	for _, x := range s.data {
		if !callback(x) {
			return
		}
	}
}
