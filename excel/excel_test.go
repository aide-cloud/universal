package excel

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

type Row struct {
	// 序号
	Index int `ex:"head:序号;index:0;"`
	// 姓名
	Name string `ex:"head:姓名;"`
	// 年龄
	Age    int      `ex:"head:年龄;index:2;"`
	Others Contents `json:"others" ex:"other:true;"`
}

func TestExcel_getTags(t *testing.T) {
	var r *[]*Row
	e, err := NewExcel("work.xlsx", "Sheet1")
	if err != nil {
		t.Fatal(err)
	}

	rType, err := e.checkTarget(r)
	if err != nil {
		t.Fatal(err)
	}

	if err = e.getTags(rType); err != nil {
		t.Fatal(err)
	}

	marshal, _ := json.Marshal(e.headers)
	t.Log(string(marshal))
}

func TestExcel_checkTarget(t *testing.T) {
	var r []*Row
	e, err := NewExcel("work.xlsx", "Sheet1")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := e.checkTarget(&r); err != nil {
		t.Fatal(err)
	}
}

func TestMarshal(t *testing.T) {
	var r []*Row
	e, err := NewExcel("work.xlsx", "Sheet1")
	if err != nil {
		t.Fatal(err)
	}

	if err := e.Marshal(&r); err != nil {
		t.Fatal(err)
	}

	marshal, _ := json.Marshal(e.headers)
	t.Log(string(marshal))

	marshal, _ = json.Marshal(r)
	t.Log(string(marshal))
}

func TestMarshalBytes(t *testing.T) {
	var r []*Row
	file, err := os.Open("work.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)

	for {
		b := make([]byte, 1024)
		n, err := file.Read(b)
		if err != nil {
			break
		}
		buf.Write(b[:n])
	}

	e, err := NewExcelWithBytes(buf.Bytes(), "Sheet1")
	if err != nil {
		t.Fatal(err)
	}

	if err := e.Marshal(&r); err != nil {
		t.Fatal(err)
	}

	marshal, _ := json.Marshal(e.headers)
	t.Log(string(marshal))

	marshal, _ = json.Marshal(r)
	t.Log(string(marshal))
}
