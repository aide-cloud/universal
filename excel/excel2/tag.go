package excel2

import (
	"errors"
	"reflect"
	"strings"
)

// 解析tag
func parseTag(tag string) (map[string]string, error) {
	tagMap := make(map[string]string)
	tags := strings.Split(tag, TagSplit)
	for _, tag := range tags {
		kv := strings.Split(tag, TagKVSplit)
		if len(kv) != 2 {
			return nil, errors.New("tag format error")
		}
		tagMap[kv[0]] = kv[1]
	}
	return tagMap, nil
}

// 获取结构体tag
func getTag(v reflect.Type) (map[string]string, error) {
	// data必须是指针类型的切片, 且切片元素必须是指针类型
	// 校验data数据类型是否正确
	if !isSliceStructPointer(v) {
		return nil, ErrDataNotSliceStructPointer
	}

	// 获取结构体
	v = v.Elem()
	// 校验data数据类型是否正确
	if !isSlice(v) {
		return nil, errors.New("data must be slice")
	}
	v = v.Elem()
	if !isPointer(v) {
		return nil, errors.New("data must be slice struct pointer")
	}
	v = v.Elem()
	if !isStruct(v) {
		return nil, errors.New("data must be slice struct pointer")
	}

	t := v
	tagMap := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(TagKey)
		if tag == "" {
			continue
		}
		tagMap[field.Name] = tag
	}

	return tagMap, nil
}

// 构建结构体属性-tag映射
func buildTagMap(v reflect.Type) (map[string]Tag, error) {
	tagMap, err := getTag(v)
	if err != nil {
		return nil, err
	}

	tagMap2 := make(map[string]Tag)
	for k, v := range tagMap {
		tag, err := parseTag(v)
		if err != nil {
			return nil, err
		}
		newTag := Tag{}
		for k2, v2 := range tag {
			switch k2 {
			case TagHead:
				headVal := v2
				newTag.Head = &headVal
			case TagIndex:
				// index属性必须是整数
				index, err := toInt(v2)
				if err != nil {
					return nil, err
				}
				indexVal := index
				newTag.Index = &indexVal
			case TagOther:
				// other属性必须是布尔字符串
				if v2 != TagTrueVal && v2 != TagFalseVal {
					return nil, errors.New("other tag value must be true or false")
				}
				otherVal := v2 == TagTrueVal
				newTag.Other = &otherVal
			}
		}
		tagMap2[k] = newTag
	}

	return tagMap2, nil
}
