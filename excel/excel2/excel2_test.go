package excel2

import (
	"encoding/json"
	"github.com/xuri/excelize/v2"
	"strconv"
	"testing"
)

func TestExcel_MarshalExcel(t *testing.T) {
	excelInstance, err := NewExcel("test2.xlsx", WithSheet("Sheet1"))
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		File     *excelize.File
		filename string
		sheets   []string
		head     map[string]Head
		data     map[string]Data
	}
	type args struct {
		data  interface{}
		sheet string
	}

	type TestStruct struct {
		Name   string `excel:"head:姓名" json:"name"`
		Age    int    `excel:"head:年龄;index:2" json:"age"`
		Index  int    `excel:"head:序号;index:0" json:"index"`
		Gender string `excel:"head:性别;index:3" json:"gender"`
	}

	var data []*TestStruct

	// {"序号", "姓名", "年龄", "性别"}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				File:     excelInstance.File,
				sheets:   excelInstance.sheets,
				head:     excelInstance.head,
				data:     excelInstance.data,
				filename: excelInstance.filename,
			},
			args: args{
				data:  &data,
				sheet: "Sheet1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Excel{
				File:     tt.fields.File,
				filename: tt.fields.filename,
				sheets:   tt.fields.sheets,
				head:     tt.fields.head,
				data:     tt.fields.data,
			}
			if err := l.MarshalExcel(tt.args.data, tt.args.sheet); (err != nil) != tt.wantErr {
				t.Errorf("MarshalExcel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		targetStr, _ := json.Marshal(tt.args.data)
		t.Log("test2.xlsx", string(targetStr))
	}
}

func TestExcel_UnmarshalExcel(t *testing.T) {
	excelInstance, err := NewExcel("test1.xlsx", WithSheet("Sheet1"))
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		File     *excelize.File
		filename string
		sheets   []string
		head     map[string]Head
		data     map[string]Data
	}
	type args struct {
		data  interface{}
		sheet string
	}

	type TestStruct struct {
		Name   string `excel:"head:姓名;index:1" json:"name"`
		Age    int    `excel:"head:年龄;index:2" json:"age"`
		Index  int    `excel:"head:序号;index:0" json:"index"`
		Gender string `excel:"head:性别;index:3" json:"gender"`
		Remark string `excel:"head:备注" json:"remark"`
	}

	var data []*TestStruct
	for i := 0; i < 10000; i++ {
		data = append(data, &TestStruct{
			Name:   "张三_" + strconv.Itoa(i),
			Age:    18 + i,
			Index:  i,
			Gender: "男",
			Remark: "备注_" + strconv.Itoa(i),
		})
	}

	// {"序号", "姓名", "年龄", "性别"}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				File:     excelInstance.File,
				sheets:   excelInstance.sheets,
				head:     excelInstance.head,
				data:     excelInstance.data,
				filename: excelInstance.filename,
			},
			args: args{
				data:  &data,
				sheet: "Sheet1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Excel{
				File:     tt.fields.File,
				filename: tt.fields.filename,
				sheets:   tt.fields.sheets,
				head:     tt.fields.head,
				data:     tt.fields.data,
			}
			if err := l.UnmarshalExcel(tt.args.data, tt.args.sheet); (err != nil) != tt.wantErr {
				t.Errorf("MarshalExcel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
