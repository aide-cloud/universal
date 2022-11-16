package basic

import "fmt"

// SliceToNumber converts a slice of type T to a slice of type Basic.
func SliceToNumber[T Number, R Number](s []T) []R {
	newSlice := make([]R, 0, len(s))
	for _, v := range s {
		newSlice = append(newSlice, R(v))
	}

	return newSlice
}

// SliceToString converts a slice of type T to a slice of type string.
func SliceToString[T any](s []T) []string {
	newSlice := make([]string, 0, len(s))
	for _, v := range s {
		newSlice = append(newSlice, fmt.Sprintf("%+v", v))
	}

	return newSlice
}
