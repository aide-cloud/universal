package assert

import (
	"reflect"
)

// IsStruct 判断参数是否为结构体或者结构体指针
func IsStruct(i any) bool {
	return isStruct(reflect.TypeOf(i))
}

// IsStructPtr 判断参数是否为结构体指针
func IsStructPtr(i any) bool {
	return isStructPtr(reflect.TypeOf(i))
}

// isStructPtr 判断参数是否为结构体指针
func isStructPtr(t reflect.Type) bool {
	if t == nil || t.Kind() != reflect.Ptr {
		return false
	}

	t = t.Elem()

	return t.Kind() == reflect.Struct
}

// isStruct 判断参数是否为结构体
func isStruct(t reflect.Type) bool {
	if t == nil || t.Kind() == reflect.Ptr {
		return false
	}

	return t.Kind() == reflect.Struct
}

// StructToMap 将结构体转换为map
func StructToMap[T any](i T) map[string]interface{} {
	// 通过反射获取类型
	t := reflect.TypeOf(i)

	if !isStruct(t) && !isStructPtr(t) {
		return nil
	}

	// 通过反射获取值
	v := reflect.ValueOf(i)

	// 创建map
	m := make(map[string]interface{})

	// 遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		// 获取字段
		field := t.Field(i)
		mapKey := field.Name
		if tag, ok := field.Tag.Lookup("json"); ok {
			mapKey = tag
		}

		// 获取字段的值
		value := v.Field(i).Interface()

		if tVal := StructToMap(value); tVal != nil {
			m[mapKey] = tVal
		} else {
			m[mapKey] = value
		}
	}

	return m
}
