package excel2

import (
	"bytes"
	"github.com/xuri/excelize/v2"
)

type (
	// ExcelOption Excel对象选项, 用于设置Excel对象的属性
	ExcelOption func(*Excel)
)

// WithSheet 设置工作表
func WithSheet(sheets ...string) ExcelOption {
	return func(e *Excel) {
		e.sheets = sheets
	}
}

// WithFilename 设置文件名
func WithFilename(filename string) ExcelOption {
	return func(e *Excel) {
		e.filename = filename
	}
}

func newExcel(ef *excelize.File) *Excel {
	return &Excel{
		sheets: make([]string, 0),
		head:   make(map[string]Head),
		data:   make(map[string]Data),
		File:   ef,
	}
}

// NewExcelWithBytes 从字节流创建一个Excel对象
func NewExcelWithBytes(data []byte, opts ...ExcelOption) (*Excel, error) {
	ef, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	excelInstance := newExcel(ef)
	for _, opt := range opts {
		opt(excelInstance)
	}

	return excelInstance, err
}

// NewExcelWithURL 从URL创建一个Excel对象
func NewExcelWithURL(url string, opts ...ExcelOption) (*Excel, error) {
	file, err := Download(url)
	if err != nil {
		return nil, err
	}

	ef, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}

	excelInstance := newExcel(ef)
	for _, opt := range opts {
		opt(excelInstance)
	}

	return excelInstance, err
}

// NewExcel 创建一个Excel对象, 根据文件名创建
func NewExcel(filename string, opts ...ExcelOption) (*Excel, error) {
	ef, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	excelInstance := newExcel(ef)
	excelInstance.filename = filename
	for _, opt := range opts {
		opt(excelInstance)
	}

	return excelInstance, err
}
