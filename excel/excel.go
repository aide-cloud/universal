package excel

import (
	"fmt"
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
		Head  string
		Index int
		Other bool
	}

	Excel struct {
		ef         *excelize.File
		rows       [][]string
		headers    map[string]*Setting
		otherKey   string
		headRowMap map[string]int
	}
)

func NewExcel(name, sheet string) (*Excel, error) {
	ef, err := excelize.OpenFile(name)
	if err != nil {
		return nil, err
	}

	rows, err := ef.GetRows(sheet)

	if len(rows) == 0 {
		return nil, fmt.Errorf("rows is empty")
	}

	headRow := rows[0]
	headRowMap := make(map[string]int)
	for index, head := range headRow {
		if _, ok := headRowMap[head]; ok {
			return nil, fmt.Errorf("head is duplicate")
		}
		headRowMap[head] = index
	}

	return &Excel{
		ef:         ef,
		rows:       rows,
		headers:    make(map[string]*Setting),
		headRowMap: headRowMap,
	}, err
}

// 解析 "head:标题;type:string;index:0;"到Setting
func (l *Excel) parseTag(tag string) *Setting {
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
	}

	return setting
}

func (l *Excel) checkTarget(target any) (reflect.Type, error) {
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

func (l *Excel) numField(target any) (reflect.Type, error) {
	return l.checkTarget(target)
}

func (l *Excel) getTags(t reflect.Type) error {
	// 获取所有字段的tag
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(TagKey)

		if tag == "" {
			continue
		}
		setting := l.parseTag(tag)
		if setting.Other {
			l.otherKey = field.Name
		}
		l.headers[field.Name] = setting
	}

	return nil
}

func (l *Excel) Marshal(target any) error {
	// 校验是不是结构体， 不同类型按照不同的方式处理， 只处理结构体切片和结构体切片指针
	rType, err := l.checkTarget(target)
	if err != nil {
		return err
	}

	// 获取所有字段的tag
	if err = l.getTags(rType); err != nil {
		return err
	}

	numFieldT := rType
	tp := reflect.TypeOf(target)
	numField := numFieldT.NumField()
	useIndexMap := make(map[int]struct{})

	for i := 1; i < len(l.rows); i++ {
		row := l.rows[i]
		tpTmp := tp.Elem().Elem()
		if tpTmp.Kind() == reflect.Ptr {
			tpTmp = tpTmp.Elem()
		}

		v := reflect.New(tpTmp)

		for j := 0; j < numField; j++ {
			field := numFieldT.Field(j)
			if setting, ok := l.headers[field.Name]; ok {
				index := setting.Index
				if index == -1 {
					index = l.headRowMap[setting.Head]
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

		if l.otherKey != "" {
			for index, val := range row {
				if _, ok := useIndexMap[index]; ok {
					continue
				}

				c := &Content{
					Key:   l.rows[0][index],
					Value: val,
				}
				v.Elem().FieldByName(l.otherKey).Set(reflect.Append(v.Elem().FieldByName(l.otherKey), reflect.ValueOf(c)))
			}
		}

		reflect.ValueOf(target).Elem().Set(reflect.Append(reflect.ValueOf(target).Elem(), v))
	}

	return nil
}
