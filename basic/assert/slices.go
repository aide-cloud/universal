package assert

import (
	"reflect"
)

// IsSlice returns true if the given value is a slice.
func IsSlice(i any) bool {
	return isSlice(reflect.TypeOf(i))
}

// IsArray returns true if the given value is an array.
func IsArray(i any) bool {
	return isArray(reflect.TypeOf(i))
}

// isArray returns true if the given value is an array.
func isArray(t reflect.Type) bool {
	if t == nil {
		return false
	}

	if t.Kind() == reflect.Ptr {
		return isArray(t.Elem())
	}

	return t.Kind() == reflect.Array
}

// isSlice returns true if the given value is a slice.
func isSlice(t reflect.Type) bool {
	if t == nil {
		return false
	}

	if t.Kind() == reflect.Ptr {
		return isSlice(t.Elem())
	}

	return t.Kind() == reflect.Slice
}
