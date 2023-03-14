package hash

import (
	"strconv"
	"testing"
)

func TestNewHashMap(t *testing.T) {
	m := NewHashMap()

	m.Add("name", "zhang.san")

	if m.Get("name") != "zhang.san" {
		t.Fatal("need zhang.san, but get ", m.Get("name"))
	}

	if m.Size() != 1 {
		t.Fatal("need 1, but get ", m.Size())
	}
}

func TestMap_Remove(t *testing.T) {
	m := NewHashMap()

	m.Add("name", "zhang.san")

	if m.Get("name") != "zhang.san" {
		t.Fatal("need zhang.san, but get ", m.Get("name"))
	}

	m.Remove("name")

	if m.Get("name") != nil {
		t.Fatal("need nil, but get ", m.Get("name"))
	}

	if m.Size() != 0 {
		t.Fatal("need 0, but get ", m.Size())
	}
}

func TestMap_IsEmpty(t *testing.T) {
	type fields struct {
		size  int
		table []*Node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test1",
			fields: fields{
				size:  0,
				table: make([]*Node, 16),
			},
			want: true,
		},
		{
			name: "test2",
			fields: fields{
				size:  1,
				table: make([]*Node, 16),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &Map{
				size:  tt.fields.size,
				table: tt.fields.table,
			}
			if got := hm.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Clear(t *testing.T) {
	m := NewHashMap()

	m.Add("name", "zhang.san")

	if m.Get("name") != "zhang.san" {
		t.Fatal("need zhang.san, but get ", m.Get("name"))
	}

	if m.Size() != 1 {
		t.Fatal("need 1, but get ", m.Size())
	}

	m.Clear()

	if m.Get("name") != nil {
		t.Fatal("need nil, but get ", m.Get("name"))
	}

	if m.Size() != 0 {
		t.Fatal("need 0, but get ", m.Size())
	}
}

func TestMap_Range(t *testing.T) {
	m := NewHashMap()

	for i := 0; i < 10; i++ {
		m.Add(strconv.Itoa(i), i)
	}

	m.Range(func(key string, value interface{}) {
		t.Log("key: ", key, ", value: ", value)
	})
}

func TestMap_Keys(t *testing.T) {
	m := NewHashMap()

	for i := 0; i < 10; i++ {
		m.Add(strconv.Itoa(i), i)
	}

	keys := m.Keys()

	if len(keys) != 10 {
		t.Fatal("need 10, but get ", len(keys))
	}

	for i := 0; i < 10; i++ {
		if keys[i] != strconv.Itoa(i) {
			t.Fatal("need ", strconv.Itoa(i), ", but get ", keys[i])
		}
	}
}

func TestMap_Values(t *testing.T) {
	m := NewHashMap()

	for i := 0; i < 10; i++ {
		m.Add(strconv.Itoa(i), i)
	}

	values := m.Values()

	if len(values) != 10 {
		t.Fatal("need 10, but get ", len(values))
	}

	for i := 0; i < 10; i++ {
		if values[i] != i {
			t.Fatal("need ", i, ", but get ", values[i])
		}
	}
}

func TestMap_ContainsKey(t *testing.T) {
	m := NewHashMap()

	for i := 0; i < 10; i++ {
		m.Add(strconv.Itoa(i), i)
	}

	if !m.ContainsKey("5") {
		t.Fatal("need true, but get false")
	}

	if m.ContainsKey("11") {
		t.Fatal("need false, but get true")
	}
}
