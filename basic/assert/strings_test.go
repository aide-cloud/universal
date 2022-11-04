package assert

import (
	"testing"
)

func TestToString(t *testing.T) {
	t.Log(ToString("test"))
	t.Log(ToString([]byte("test")))
	t.Log(ToString(nil))
	t.Log(ToString(1))
	t.Log(ToString(1.1))
	t.Log(ToString(true))
	t.Log(ToString(false))
	t.Log(ToString([]int{1, 2, 3}))
	t.Log(ToString(map[string]int{"a": 1, "b": 2, "c": 3}))
	t.Log(ToString(map[int]string{1: "a", 2: "b", 3: "c"}))
	t.Log(ToString(map[int]int{1: 1, 2: 2, 3: 3}))

	var a interface{}
	t.Log(ToString(a))

	var b *int
	t.Log(ToString(&b))
	b = new(int)
	t.Log(ToString(&b))

}
