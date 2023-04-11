package excel

import (
	"bytes"
	"encoding/json"
	"os"
	"strconv"
	"testing"
)

type Row struct {
	// 序号
	Index int `ex:"head:序号;index:0;"`
	// 姓名
	Name string `ex:"head:姓名;index:2;"`
	// 年龄
	Age    int      `ex:"head:年龄;index:3;"`
	Others Contents `json:"others" ex:"other:true;"`
}

func TestExcel_checkTarget(t *testing.T) {
	var r []*Row
	if _, err := checkTarget(&r); err != nil {
		t.Fatal(err)
	}
}

func TestExcel_getTags(t *testing.T) {
	var r *[]*Row
	rType, err := checkTarget(r)
	if err != nil {
		t.Fatal(err)
	}

	headers, otherKey := getTags(rType)
	t.Log("otherKey:", otherKey)

	marshal, _ := json.Marshal(headers)
	t.Log(string(marshal))
}

func TestMarshal(t *testing.T) {
	var r []*Row
	e, err := NewExcel("work.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	if err := e.Marshal(&r, "Sheet1"); err != nil {
		t.Fatal(err)
	}

	marshal, _ := json.Marshal(r)
	t.Log(string(marshal))
}

func TestMarshalBytes(t *testing.T) {
	var r []*Row
	file, err := os.Open("work.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			t.Fatal(err)
		}
	}(file)

	buf := bytes.NewBuffer(nil)

	for {
		b := make([]byte, 1024)
		n, err := file.Read(b)
		if err != nil {
			break
		}
		buf.Write(b[:n])
	}

	e, err := NewExcelWithBytes(buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	if err := e.Marshal(&r, "Sheet1"); err != nil {
		t.Fatal(err)
	}

	marshal, _ := json.Marshal(r)
	t.Log(string(marshal))
}

func TestUnmarshal(t *testing.T) {
	var r []*Row

	for i := 1; i < 10; i++ {
		r = append(r, &Row{
			Index: i,
			Name:  "name_" + strconv.Itoa(i),
			Age:   i,
			Others: []*Content{
				{
					Key:   "key",
					Value: "value" + strconv.Itoa(i),
				},
				{
					Key:   "key_1",
					Value: "value_" + strconv.Itoa(i),
				},
				{
					Key:   "key_2",
					Value: "value_2_" + strconv.Itoa(i),
				},
			},
		})
	}

	fileName, sheet := "test.xlsx", "Sheet1"

	e, err := NewExcel(fileName)
	if err != nil {
		t.Fatal(err)
	}

	if err := e.Unmarshal(&r, sheet); err != nil {
		t.Fatal(err)
	}
}
