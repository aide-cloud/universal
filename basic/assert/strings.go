package assert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unicode/utf8"
)

// ToString converts a value to a string.
func ToString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	case fmt.Stringer:
		return val.String()
	case error:
		return val.Error()
	case nil, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
		return fmt.Sprintf("%v", val)
	default:
		marshal, err := json.Marshal(val)
		if err == nil {
			return string(marshal)
		}
		return fmt.Sprintf("%v", val)
	}
}

// Count counts the number of elements in a slice.
func Count(target any) int {
	count := 0
	switch val := target.(type) {
	case string:
		count = utf8.RuneCountInString(val)

	default:
		// 反射获取数据类型
		t := reflect.TypeOf(val)
		switch t.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map:
			count = reflect.ValueOf(val).Len()
		default:
			count = 0
		}
		count = 0
	}
	return count
}
