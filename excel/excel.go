package excel

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	TagKey     = "ex"    // Tag属性名
	TagHead    = "head"  // 是否为头属性
	TagIndex   = "index" // 是否为索引属性
	TagOther   = "other" // 是否为其他属性
	TagTrueVal = "true"  // Tag属性值为true
	TagSplit   = ";"     // Tag属性分隔符
	TagKVSplit = ":"     // Tag属性键值分隔符
)

type (
	Content struct {
		Key   string
		Value string
	}

	Contents []*Content

	Setting struct {
		FieldName string
		Head      string
		Index     int
		Other     bool
	}

	Excel struct {
		*excelize.File
	}
)

// NewExcelWithBytes 从字节流创建一个Excel对象
func NewExcelWithBytes(data []byte) (*Excel, error) {
	ef, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return &Excel{ef}, err
}

// NewExcel 创建一个Excel对象, 如果文件不存在则创建
func NewExcel(filename string) (*Excel, error) {
	if err := createFile(filename); err != nil {
		return nil, err
	}
	ef, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	return &Excel{ef}, err
}

// CreateExcelFile 创建一个新的excel文件
func (l *Excel) CreateExcelFile(fileName, sheet string) error {
	_, err := newExcellingFile(fileName, sheet)
	if err != nil {
		return err
	}

	return nil
}

// Marshal 将excel转换为结构体切片
func (l *Excel) Marshal(target any, sheet string) error {
	// 校验是不是结构体， 不同类型按照不同的方式处理， 只处理结构体切片和结构体切片指针
	rType, err := checkTarget(target)
	if err != nil {
		return err
	}

	// 获取所有字段的tag
	headers, otherKey := getTags(rType)
	rows, err := l.GetRows(sheet)
	if err != nil {
		return err
	}
	headRowMap := make(map[string]int)
	if len(rows) > 0 {
		headRow := rows[0]
		for index, head := range headRow {
			if _, ok := headRowMap[head]; ok {
				return fmt.Errorf("head is duplicate")
			}
			headRowMap[head] = index
		}
	}

	numFieldT := rType
	tp := reflect.TypeOf(target)
	numField := numFieldT.NumField()
	useIndexMap := make(map[int]struct{})
	tpTmp := tp.Elem().Elem()
	if tpTmp.Kind() == reflect.Ptr {
		tpTmp = tpTmp.Elem()
	}

	for i := 1; i < len(rows); i++ {
		row := rows[i]

		v := reflect.New(tpTmp)

		for j := 0; j < numField; j++ {
			field := numFieldT.Field(j)
			if setting, ok := headers[field.Name]; ok {
				index := setting.Index
				if index == -1 {
					index = headRowMap[setting.Head]
				}

				if setting.Index >= len(row) || setting.Other {
					continue
				}

				switch field.Type.Kind() {
				case reflect.String:
					v.Elem().Field(j).SetString(row[index])
					useIndexMap[index] = struct{}{}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					value, _ := strconv.ParseInt(row[index], 10, 64)
					v.Elem().Field(j).SetInt(value)
					useIndexMap[index] = struct{}{}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					value, _ := strconv.ParseUint(row[index], 10, 64)
					v.Elem().Field(j).SetUint(value)
					useIndexMap[index] = struct{}{}
				case reflect.Float32, reflect.Float64:
					value, _ := strconv.ParseFloat(row[index], 64)
					v.Elem().Field(j).SetFloat(value)
					useIndexMap[index] = struct{}{}
				case reflect.Bool:
					value, _ := strconv.ParseBool(row[index])
					v.Elem().Field(j).SetBool(value)
					useIndexMap[index] = struct{}{}
				}
			}
		}

		if otherKey != "" {
			for index, val := range row {
				if _, ok := useIndexMap[index]; ok {
					continue
				}

				c := &Content{
					Key:   rows[0][index],
					Value: val,
				}
				v.Elem().FieldByName(otherKey).Set(reflect.Append(v.Elem().FieldByName(otherKey), reflect.ValueOf(c)))
			}
		}

		reflect.ValueOf(target).Elem().Set(reflect.Append(reflect.ValueOf(target).Elem(), v))
	}

	return nil
}

// Unmarshal 从excel中读取数据，写入target
func (l *Excel) Unmarshal(target any, sheet string) error {
	// 把target的数据写入excel
	if err := l.writeExcel(target, sheet); err != nil {
		return err
	}

	return nil
}

// 解析 "head:标题;type:string;index:0;"到Setting
func parseTag(tag string) *Setting {
	setting := &Setting{
		Head:  "",
		Index: -1,
		Other: false,
	}
	for _, v := range strings.Split(tag, TagSplit) {
		kv := strings.Split(v, TagKVSplit)
		switch kv[0] {
		case TagHead:
			setting.Head = kv[1]
		case TagIndex:
			setting.Index, _ = strconv.Atoi(kv[1])
		case TagOther:
			setting.Other = kv[1] == TagTrueVal
		}
		setting.FieldName = strings.ToLower(setting.Head)
	}

	return setting
}

// checkTarget 校验目标类型. 目标类型必须是一个结构体切片或者结构体切片指针
func checkTarget(target any) (reflect.Type, error) {
	// 校验是不是结构体， 不同类型按照不同的方式处理， 只处理结构体切片和结构体切片指针
	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	} else {
		return nil, fmt.Errorf("must be a slice or slice pointer")
	}
	if t.Kind() != reflect.Slice {
		return nil, fmt.Errorf("must be a slice or slice pointer")
	}

	t = t.Elem()

	if t.Kind() != reflect.Ptr {
		// 切片必须是一个指针的切片， 不能是一个结构体的切片
		return nil, fmt.Errorf("must be a slice pointer")
	}

	return t.Elem(), nil
}

// getTags 获取所有字段的tag
func getTags(t reflect.Type) (map[string]*Setting, string) {
	headers := make(map[string]*Setting)
	otherKey := ""
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// index标记， 用于标记index是否重复
	indexMap := make(map[int]string)

	// 获取所有字段的tag
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(TagKey)

		if tag == "" {
			continue
		}
		setting := parseTag(tag)

		// 如果index重复，则报错
		if fieldName, ok := indexMap[setting.Index]; ok {
			panic(fmt.Sprintf("index %d is duplicate, field name's '%s' and '%s'", setting.Index, fieldName, field.Name))
		}

		if setting.Other {
			otherKey = field.Name
		}

		headers[field.Name] = setting

		indexMap[setting.Index] = field.Name
	}

	return headers, otherKey
}

// newExcellingFile 创建一个新的excel文件
func newExcellingFile(fileName, sheet string) (*excelize.File, error) {
	// 判断文件是否存在， 不存在则创建
	if err := createFile(fileName); err != nil {
		return nil, err
	}

	// 判断是否有表头， 没有则根据target的字段名生成表头，写sheet
	// 读取excel
	ef, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	// 如果没有sheet，则创建sheet
	sheetIndex, err := ef.GetSheetIndex(sheet)
	if err != nil {
		if err != nil || sheetIndex == -1 {
			_, err = ef.NewSheet(sheet)
			if err != nil {
				return nil, err
			}
		}
	}

	return ef, nil
}

// getHeaderLen 获取表头长度
func getHeaderLen(headers map[string]*Setting, othersKey string) int {
	headerRowLen := len(headers)
	if o, ok := headers[othersKey]; ok && o.Other {
		headerRowLen--
	}

	for _, v := range headers {
		if v.Index >= headerRowLen {
			headerRowLen = v.Index + 1
		}
	}

	return headerRowLen
}

// getTargetVElem 获取target的数据
func getTargetVElem(target any) (reflect.Value, error) {
	// 获取target的数据
	v := reflect.ValueOf(target)
	// panic: reflect: call of reflect.Value.Len on ptr Value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Slice {
		return reflect.Value{}, errors.New("target must be slice")
	}

	return v, nil
}

// getHeaderRow 获取表头
func getHeaderRow(item reflect.Value, headers map[string]*Setting, headerRowLen int, othersKey string) []interface{} {
	// 设置表头
	headerRow := make([]interface{}, headerRowLen)
	for _, setting := range headers {
		insertIndex := setting.Index
		if insertIndex == -1 || setting.Other {
			if item.Kind() == reflect.Ptr {
				item = item.Elem()
			}
			// 设置表头， 根据Others的列表遍历插入other的key值作为表头
			others := item.FieldByName(othersKey)
			if others.Kind() == reflect.Ptr {
				others = others.Elem()
			}
			switch others.Kind() {
			case reflect.Slice:
				for j := 0; j < others.Len(); j++ {
					oItem := others.Index(j)
					if oItem.Kind() == reflect.Ptr {
						oItem = oItem.Elem()
					}
					headerRow = append(headerRow, oItem.FieldByName("Key").Interface())
				}
			}
			continue
		}
		if insertIndex < 0 || insertIndex >= len(headerRow) {
			insertIndex = 0
		}

		headerRow[insertIndex] = setting.Head
	}

	return headerRow
}

// getDataRow 获取数据行
func getDataRow(sliceVal reflect.Value, headers map[string]*Setting, headerRowLen int, othersKey string) [][]interface{} {
	rows := make([][]interface{}, sliceVal.Len())
	for i := 0; i < sliceVal.Len(); i++ {
		row := make([]interface{}, headerRowLen)
		item := sliceVal.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		for j := 0; j < item.NumField(); j++ {
			field := item.Field(j)
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			// 根据字段名插入数据
			if setting, ok := headers[item.Type().Field(j).Name]; ok {
				insertIndex := setting.Index
				if setting.Other {
					others := item.FieldByName(othersKey)
					// 设置数据
					switch others.Kind() {
					case reflect.Slice:
						for oIndex := 0; oIndex < others.Len(); oIndex++ {
							oItem := others.Index(oIndex)
							if oItem.Kind() == reflect.Ptr {
								oItem = oItem.Elem()
							}
							row = append(row, oItem.FieldByName("Value").Interface())
						}
					}
				} else {
					if insertIndex < 0 {
						insertIndex = 0
					}

					row[insertIndex] = field.Interface()
				}
			}
		}

		rows[i] = row
	}

	return rows
}

// writeExcel 把target的数据写入excel
func (l *Excel) writeExcel(target any, sheet string) error {
	// 把target的数据写入excel
	rType, err := checkTarget(target)
	if err != nil {
		return err
	}

	// 获取所有字段的tag
	headers, othersKey := getTags(rType)

	// 获取target的数据
	v, err := getTargetVElem(target)
	if err != nil {
		return err
	}

	headerRowLen := getHeaderLen(headers, othersKey)

	rows := make([][]interface{}, 0, v.Len())
	if v.Len() > 0 {
		// 设置表头
		headerRow := getHeaderRow(v.Index(0), headers, headerRowLen, othersKey)
		rows = append(rows, headerRow)
		// 设置数据
		rows = append(rows, getDataRow(v, headers, headerRowLen, othersKey)...)
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

	return nil
}

// createFile 创建文件
func createFile(fileName string) error {
	// 判断文件是否存在， 不存在则创建
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			// 创建一个xlsx文件
			f := excelize.NewFile()
			// 创建一个工作表
			sheet, err := f.NewSheet("Sheet1")
			if err != nil {
				return err
			}

			// 设置工作表为默认工作表
			f.SetActiveSheet(sheet)
			// 保存文件
			if err := f.SaveAs(fileName); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
