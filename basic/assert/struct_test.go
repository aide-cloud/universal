package assert

import (
	"encoding/json"
	"fmt"
	"testing"
)

type (
	S struct {
		A          *int     `json:"a"`
		B          string   `json:"b_xxx"`
		C          bool     `json:"c"`
		D          []int    `json:"d"`
		DDStr      []string `json:"dd_str"`
		StructList []TMP    `json:"struct_list"`
		Str        struct {
			A1 int    `json:"a_1"`
			B1 string `json:"b_1"`
			C1 bool   `json:"c_1"`
		} `json:"str"`
	}
	TMP struct {
		A int `json:"a"`
	}
)

func TestIsStruct(t *testing.T) {
	var i interface{}
	var s S
	if IsStruct(1) {
		t.Error("1 is not struct")
	}

	if IsStruct(i) {
		t.Error("nil is not struct")
	}

	if !IsStruct(s) {
		t.Error("struct is struct")
	}

	if IsStruct(new(S)) {
		t.Error("struct ptr not is struct")
	}
}

func TestIsStructPtr(t *testing.T) {
	var i interface{}
	var s S
	if IsStructPtr(1) {
		t.Error("1 is not struct")
	}

	if IsStructPtr(i) {
		t.Error("nil is not struct")
	}

	if IsStructPtr(s) {
		t.Error("struct not is struct ptr")
	}

	if !IsStructPtr(new(S)) {
		t.Error("struct ptr not is struct")
	}
}

func TestStructToMap(t *testing.T) {
	var s S
	s.D = []int{1, 2, 3}
	s.DDStr = []string{"1", "2", "3"}
	s.StructList = []TMP{{A: 1}, {A: 2}}
	m := StructToMap(s)
	if m == nil {
		t.Error("struct to map is nil")
	}
	t.Log(m)
	fmt.Printf("%#v\n", m)
	marshal, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(marshal))

	//m = StructToMap(new(S))
	//t.Log(m)
	//if m == nil {
	//	t.Error("struct ptr to map is nil")
	//}
	//
	//var i interface{}
	//m = StructToMap(i)
	//if m != nil {
	//	t.Error("interface to map is not nil")
	//}
	//
	//m = StructToMap(1)
	//if m != nil {
	//	t.Error("int to map is not nil")
	//}
	//
	//m = StructToMap("1")
	//if m != nil {
	//	t.Error("string to map is not nil")
	//}
	//
	//m = StructToMap(true)
	//if m != nil {
	//	t.Error("bool to map is not nil")
	//}
}
