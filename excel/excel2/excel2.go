package excel2

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
)

const (
	TagKey      = "excel" // Tag属性名
	TagHead     = "head"  // 是否为头属性
	TagIndex    = "index" // 是否为索引属性
	TagOther    = "other" // 是否为其他属性
	TagTrueVal  = "true"  // Tag属性值为true
	TagFalseVal = "false" // Tag属性值为false
	TagSplit    = ";"     // Tag属性分隔符
	TagKVSplit  = ":"     // Tag属性键值分隔符
)

type (
	Head = []string
	Data = [][]string

	Excel struct {
		*excelize.File
		filename string
		sheets   []string
		head     map[string]Head // sheetName -> head
		data     map[string]Data // sheetName -> data
	}

	Tag struct {
		Head  *string
		Index *int
		Other *bool
	}

	// ExcelMarshal 解析和映射的接口
	ExcelMarshal interface {
		// UnmarshalExcel 解析
		UnmarshalExcel(data interface{}, sheet string) error
		// MarshalExcel 映射
		MarshalExcel(data interface{}, sheet string) error
	}
)

var (
	ErrDataNotSliceStructPointer = errors.New("data必须是指针类型的切片, 且切片元素必须是结构体指针类型")
)

// UnmarshalExcel 解析, 把target数据写到Excel对象中
func (l *Excel) UnmarshalExcel(target interface{}, sheet string) error {
	// data必须是指针类型的切片, 且切片元素必须是指针类型
	// 校验data数据类型是否正确
	t := reflect.TypeOf(target)
	if !isSliceStructPointer(t) {
		return ErrDataNotSliceStructPointer
	}

	// 获取data数据的类型的tag信息
	dataTagMap, err := buildTagMap(t)
	if err != nil {
		return err
	}

	head, err := l.getHead(dataTagMap)
	if err != nil {
		return err
	}

	rows, err := l.structDataToSliceData(target, head, dataTagMap)
	if err != nil {
		return err
	}

	// 批量写入excel
	writer, err := l.NewStreamWriter(sheet)
	if err != nil {
		return err
	}

	// 写入数据
	for i, row := range rows {
		if err = writer.SetRow(fmt.Sprintf("A%d", i+1), row); err != nil {
			return err
		}
	}

	if err = writer.Flush(); err != nil {
		return err
	}

	if err = l.Save(); err != nil {
		return err
	}

	return err
}

// MarshalExcel 映射, 把Excel对象中的数据写到target数据中
func (l *Excel) MarshalExcel(target interface{}, sheet string) error {
	// data必须是指针类型的切片, 且切片元素必须是指针类型
	t := reflect.TypeOf(target)
	if !isSliceStructPointer(t) {
		return ErrDataNotSliceStructPointer
	}

	// 获取data数据的类型的tag信息
	dataTagMap, err := buildTagMap(t)
	if err != nil {
		return err
	}

	if err := l.setSheetDataWithHeads(); err != nil {
		return err
	}

	sheetData, found := l.data[sheet]
	if !found {
		return errors.New(fmt.Sprintf("sheet %s not found", sheet))
	}

	if err := l.sliceDataToStructData(target, sheetData, dataTagMap, sheet); err != nil {
		return err
	}

	return nil
}

var _ ExcelMarshal = (*Excel)(nil)
