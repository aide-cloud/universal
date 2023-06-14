package excel2

import (
	"github.com/xuri/excelize/v2"
	"reflect"
	"testing"
)

func TestExcel_GetSheetNoneData(t *testing.T) {
	excelInstance, err := NewExcel("test1.xlsx")
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
		sheet string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				File: excelInstance.File,
			},
			args: args{
				sheet: "Sheet1",
			},
			want:    [][]string{},
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
			got, err := l.GetSheetData(tt.args.sheet)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSheetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSheetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExcel_GetSheetHaveData(t *testing.T) {
	excelInstance, err := NewExcel("test2.xlsx")
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
		sheet string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				File: excelInstance.File,
			},
			args: args{
				sheet: "Sheet1",
			},
			want: [][]string{
				{"序号", "姓名", "年龄", "性别"},
				{"1", "张三", "18", "男"},
				{"2", "李四", "19", "女"},
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
			got, err := l.GetSheetData(tt.args.sheet)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSheetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSheetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExcel_setSheetDataWithHead(t *testing.T) {
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
		sheet string
	}
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
			if err := l.setSheetDataWithHead(tt.args.sheet); (err != nil) != tt.wantErr {
				t.Errorf("setSheetDataWithHead() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExcel_setSheetDataWithHeads(t *testing.T) {
	excelInstance, err := NewExcel("test2.xlsx", WithSheet("Sheet1", "Sheet2"))
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
	tests := []struct {
		name   string
		fields fields
		want   []error
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
			want: []error{nil, nil},
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
			if got := l.setSheetDataWithHeads(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setSheetDataWithHeads() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExcel_GetHead(t *testing.T) {
	excelInstance, err := NewExcel("test2.xlsx", WithSheet("Sheet1", "Sheet2"))
	if err != nil {
		t.Fatal(err)
	}

	if err := excelInstance.setSheetDataWithHeads(); err != nil {
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
		sheet string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Head
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
				sheet: "Sheet1",
			},
			want:    Head{"序号", "姓名", "年龄", "性别"},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				File:     excelInstance.File,
				sheets:   excelInstance.sheets,
				head:     excelInstance.head,
				data:     excelInstance.data,
				filename: excelInstance.filename,
			},
			args: args{
				sheet: "Sheet2",
			},
			want:    nil,
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
			got, err := l.GetHead(tt.args.sheet)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHead() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExcel_GetData(t *testing.T) {
	excelInstance, err := NewExcel("test2.xlsx", WithSheet("Sheet1", "Sheet2"))
	if err != nil {
		t.Fatal(err)
	}
	if err := excelInstance.setSheetDataWithHeads(); err != nil {
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
		sheet string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Data
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
				sheet: "Sheet1",
			},
			want:    Data{{"1", "张三", "18", "男"}, {"2", "李四", "19", "女"}},
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				File:     excelInstance.File,
				sheets:   excelInstance.sheets,
				head:     excelInstance.head,
				data:     excelInstance.data,
				filename: excelInstance.filename,
			},
			args: args{
				sheet: "Sheet2",
			},
			want:    nil,
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
			got, err := l.GetData(tt.args.sheet)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
