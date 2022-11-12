package utils

type Set[T comparable] struct {
	data []T
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: []T{},
	}
}

func (set *Set[T]) Add(el T) {
	set.data = append(set.data, el)
}

func (set *Set[T]) Remove(el T) {
	for i, v := range set.data {
		if v == el {
			s := set.data
			s[i] = s[len(s)-1]
			set.data = s[:len(s)-1]
		}
	}
}

func (set *Set[T]) Contains(el T) bool {
	for _, v := range set.data {
		if v == el {
			return true
		}
	}
	return false
}

func (set *Set[T]) Values() []T {
	return set.data
}

func (set *Set[T]) Size() int {
	return len(set.data)
}

func ArraytoSet[T comparable](array []T) *Set[T] {
	set := NewSet[T]()
	for _, v := range array {
		set.Add(v)
	}
	return set
}
