package excel2

import (
	"errors"
	"fmt"
	"github.com/aide-cloud/universal/base/slices"
	"reflect"
)

// GetSheetData 获取表格数据
func (l *Excel) GetSheetData(sheet string) ([][]string, error) {
	return l.GetRows(sheet)
}

// setSheetDataWithHead 根据sheet名获取表格数据, 并设置表头和数据
func (l *Excel) setSheetDataWithHead(sheet string) error {
	data, err := l.GetSheetData(sheet)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		if l.head == nil {
			l.head = make(map[string]Head)
		}
		if _, found := l.head[sheet]; !found {
			l.head[sheet] = make(Head, 0, len(data[0]))
		}
		l.head[sheet] = data[0]

		if len(data) > 1 {
			if l.data == nil {
				l.data = make(map[string]Data)
			}
			if _, found := l.data[sheet]; !found {
				l.data[sheet] = make(Data, 0, len(data)-1)
			}
			l.data[sheet] = data[1:]
		}
	}

	return nil
}

// 批量设置表格数据
func (l *Excel) setSheetDataWithHeads() error {
	for _, sheet := range l.sheets {
		if err := l.setSheetDataWithHead(sheet); err != nil {
			return err
		}
	}

	return nil
}

// GetHead 获取head
func (l *Excel) GetHead(sheet string) (Head, error) {
	if l.head == nil {
		if err := l.setSheetDataWithHeads(); err != nil {
			return nil, errors.New("sheet not found")
		}
	}

	return l.head[sheet], nil
}

// GetData 获取data
func (l *Excel) GetData(sheet string) (Data, error) {
	if l.data == nil {
		if err := l.setSheetDataWithHeads(); err != nil {
			return nil, errors.New("sheet not found")
		}
	}

	return l.data[sheet], nil
}

// 获取表头
func (l *Excel) getHead(dataTagMap map[string]Tag) ([]interface{}, error) {
	head := make([]interface{}, len(dataTagMap))
	noIndexHead := make([]interface{}, 0, len(dataTagMap))
	for k, v := range dataTagMap {
		if v.Index != nil && v.Head != nil {
			if *v.Index >= len(head) {
				return nil, fmt.Errorf("index: %d, out of range: %d", *v.Index, len(head))
			}
			headName := k
			if *v.Head != "" {
				headName = *v.Head
			}
			head[*v.Index] = headName
			continue
		}

		if v.Head != nil {
			noIndexHead = append(noIndexHead, *v.Head)
		}
	}

	for i, v := range head {
		if len(noIndexHead) == 0 {
			break
		}
		if v == nil {
			head[i] = noIndexHead[0]
			noIndexHead = noIndexHead[1:]
		}
	}

	return head, nil
}

// 把结构体数据转成二维数组
func (l *Excel) structDataToSliceData(target interface{}, head []interface{}, dataTagMap map[string]Tag) ([][]interface{}, error) {
	values := reflect.ValueOf(target).Elem()

	rows := make([][]interface{}, 0, values.Len())

	// 遍历target, 把数据从target中写到放到rows中
	for i := 0; i < values.Len(); i++ {
		length := values.Index(i).Elem().NumField()
		row := make([]interface{}, len(head))
		for j := 0; j < length; j++ {
			//row = append(row, values.Index(i).Elem().Field(j).Interface())
			field := values.Index(i).Elem().Type().Field(j)
			index := -1
			for k, v := range dataTagMap {
				if k == field.Name {
					if v.Index != nil {
						index = *v.Index
						break
					}
					headName := k
					if v.Head != nil {
						headName = *v.Head
					}
					for i, v := range head {
						val := fmt.Sprintf("%v", v)
						if val == headName {
							index = i
							break
						}
					}
				}
			}

			if index == -1 || index >= length {
				continue
			}

			row[index] = fmt.Sprintf("%v", values.Index(i).Elem().Field(j).Interface())
		}
		rows = append(rows, row)
	}
	rows = append([][]interface{}{head}, rows...)
	return rows, nil
}

// 把二维数组数据转成结构体数据
func (l *Excel) sliceDataToStructData(target interface{}, sheetData Data, dataTagMap map[string]Tag, sheet string) error {
	// 把数据从Excel对象中写到data数据中
	values := reflect.TypeOf(target)
	numField := values.Elem().Elem().Elem().NumField()
	for _, data := range sheetData {
		// 把数据写到data数据中
		v := reflect.New(values.Elem().Elem().Elem())
		// 遍历字段
		for i := 0; i < numField; i++ {
			field := values.Elem().Elem().Elem().Field(i)
			tag := field.Tag.Get(TagKey)
			if tag == "" {
				continue
			}
			tagMap, found := dataTagMap[field.Name]
			if !found {
				continue
			}

			headIndex := -1
			// 字符串类型
			if tagMap.Head != nil {
				// 头属性
				headIndex = slices.Index(l.head[sheet], *tagMap.Head)
			}
			if tagMap.Index != nil {
				// 索引属性
				headIndex = *tagMap.Index
			}

			switch field.Type.Kind() {
			case reflect.String:
				v.Elem().Field(i).SetString(data[headIndex])
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				val, _ := toInt(data[headIndex])
				v.Elem().Field(i).SetInt(int64(val))
			case reflect.Float32, reflect.Float64:
				val, _ := toFloat(data[headIndex])
				v.Elem().Field(i).SetFloat(val)
			}
		}

		// 把v按照data数据的类型写到data数据中
		reflect.ValueOf(target).Elem().Set(reflect.Append(reflect.ValueOf(target).Elem(), v))
	}

	return nil
}
