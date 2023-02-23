package slices

import (
	"sort"
)

// Index returns the index of the first occurrence of the specified element in the specified slice.
// If the slice does not contain the element, it returns -1.
func Index[T comparable](a []T, x T) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

// IndexStruct returns the index of the first occurrence of the specified element in the specified slice.
// If the slice does not contain the element, it returns -1.
func IndexStruct[T comparable](a []T, f func(T) bool) int {
	for i, n := range a {
		if f(n) {
			return i
		}
	}
	return -1
}

// Contains returns true if the specified slice contains the specified element.
func Contains[T comparable](a []T, x T) bool {
	return Index(a, x) != -1
}

// ContainsStruct returns true if the specified slice contains the specified element.
func ContainsStruct[T comparable](a []T, f func(T) bool) bool {
	return IndexStruct(a, f) != -1
}

// Filter returns a new slice containing all elements of the specified slice that satisfy the specified predicate.
func Filter[T any](a []T, predicate func(T) bool) []T {
	var b []T
	for _, x := range a {
		if predicate(x) {
			b = append(b, x)
		}
	}
	return b
}

// Map returns a new slice containing the results of applying the specified function to each element of the specified slice.
func Map[T, U any](a []T, f func(T) U) []U {
	var b []U
	for _, x := range a {
		b = append(b, f(x))
	}
	return b
}

// To returns a new slice containing the results of applying the specified function to each element of the specified slice.
func To[T, U any](a []T, f func(T) U) []U {
	return Map(a, f)
}

// Sort returns a new slice containing the elements of the specified slice in ascending order.
func Sort[T any](a []T, less func(i, j T) bool) []T {
	sort.Slice(a, func(i, j int) bool {
		return less(a[i], a[j])
	})
	return a
}

// Reverse returns a new slice containing the elements of the specified slice in reverse order.
func Reverse[T any](a []T) []T {
	b := make([]T, len(a))
	for i, j := 0, len(a)-1; i < len(a); i, j = i+1, j-1 {
		b[i] = a[j]
	}
	return b
}

// Unique returns a new slice containing the elements of the specified slice with duplicates removed.
func Unique[T comparable](a []T) []T {
	b := make([]T, 0, len(a)/2)
	tmp := make(map[T]struct{}, len(a)/2)
	for _, x := range a {
		if _, ok := tmp[x]; !ok {
			tmp[x] = struct{}{}
			b = append(b, x)
		}
	}

	return b
}

type Uniquer interface {
	UniqueKey() string
}

// UniqueStruct returns a new slice containing the elements of the specified slice with duplicates removed.
func UniqueStruct[T Uniquer](a []T) []T {
	b := make([]T, 0, len(a)/2)
	tmp := make(map[string]struct{}, len(a)/2)
	for _, x := range a {
		if _, ok := tmp[x.UniqueKey()]; !ok {
			tmp[x.UniqueKey()] = struct{}{}
			b = append(b, x)
		}
	}

	return b
}
