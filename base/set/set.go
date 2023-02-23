package set

type Set[T comparable] map[T]struct{}

// New returns a new set with the given items.
func New[T comparable](items ...T) Set[T] {
	s := Set[T](make(map[T]struct{}))
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

// Add adds an item to the set.
func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Remove removes an item from the set.
func (s Set[T]) Remove(item T) {
	delete(s, item)
}

// Has returns true if the item is in the set.
func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

// Len returns the number of items in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Clear removes all items from the set.
func (s *Set[T]) Clear() {
	*s = make(map[T]struct{})
}

// IsEmpty returns true if the set is empty.
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// ToSlice returns a slice containing all items in the set.
func (s Set[T]) ToSlice() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

// ForEach iterates over the set and calls the given function for each item.
func (s Set[T]) ForEach(f func(T)) {
	for item := range s {
		f(item)
	}
}

// Map returns a new set with the result of applying the given function to each item.
func (s Set[T]) Map(f func(T) T) Set[T] {
	r := New[T]()
	for item := range s {
		r.Add(f(item))
	}
	return r
}

// Filter returns a new set with all items that satisfy the given predicate.
func (s Set[T]) Filter(f func(T) bool) Set[T] {
	r := New[T]()
	for item := range s {
		if f(item) {
			r.Add(item)
		}
	}
	return r
}
