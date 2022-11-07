package assert

import (
	"github.com/aide-cloud/universal/basic"
	"testing"
)

func TestToString(t *testing.T) {
	t.Log(basic.ToString("test"))
	t.Log(basic.ToString([]byte("test")))
	t.Log(basic.ToString(nil))
	t.Log(basic.ToString(1))
	t.Log(basic.ToString(1.1))
	t.Log(basic.ToString(true))
	t.Log(basic.ToString(false))
	t.Log(basic.ToString([]int{1, 2, 3}))
	t.Log(basic.ToString(map[string]int{"a": 1, "b": 2, "c": 3}))
	t.Log(basic.ToString(map[int]string{1: "a", 2: "b", 3: "c"}))
	t.Log(basic.ToString(map[int]int{1: 1, 2: 2, 3: 3}))

	var a interface{}
	t.Log(basic.ToString(a))

	var b *int
	t.Log(basic.ToString(&b))
	b = new(int)
	t.Log(basic.ToString(&b))

}

func TestCount(t *testing.T) {
	//t.Log(Count("hello"))
	//t.Log(Count([]int{1, 2, 3}))
	//t.Log(Count(map[string]int{"a": 1, "b": 2, "c": 3}))
	//t.Log(Count(map[int]string{1: "a", 2: "b", 3: "c"}))
	//t.Log(Count(map[int]int{1: 1, 2: 2, 3: 3}))
	//t.Log(Count([]string{"a", "b", "c"}))
	//t.Log(Count("你好，世界"))
	//t.Log(Count([]struct{}{{}, {}, {}}))
	//t.Log(Count(struct{}{}))
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	t.Log(len(ch))
	t.Log(basic.Count(ch))
}
