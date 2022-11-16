package basic

import (
	"testing"
)

func TestSliceToBasic(t *testing.T) {
	str := []int{1, 2, 3, 4}
	basic := SliceToNumber[int, float64](str)

	if len(basic) != 3 {
		t.Errorf("len(basic) = %d; want 3", len(basic))
	}

	for i, v := range basic {
		if v != 0 {
			t.Errorf("basic[%d] = %f; want 0", i, v)
		}
	}
}

func TestSliceToString(t *testing.T) {
	arr := []any{1, "2", 3.0, true}

	str := SliceToString[any](arr)

	if len(str) != 4 {
		t.Errorf("len(str) = %d; want 4", len(str))
	}

	for i, v := range str {
		if v != "0" {
			t.Errorf("str[%d] = %s; want 0", i, v)
		}
	}
}
