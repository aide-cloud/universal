package assert

import (
	"testing"
)

func TestIsArray(t *testing.T) {
	var arr [3]int
	var slice []int

	if !IsArray(arr) {
		t.Error("IsArray([]int{}) should be true")
	}

	if IsArray(slice) {
		t.Error("IsArray([3]int{}) should be false")
	}
}

func TestIsSlice(t *testing.T) {
	var slice []int
	var arr [3]int

	if !IsSlice(slice) {
		t.Error("IsSlice([]int{}) should be true")
	}

	if IsSlice(arr) {
		t.Error("IsSlice([3]int{}) should be false")
	}
}
