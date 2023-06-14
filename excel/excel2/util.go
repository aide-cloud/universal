package excel2

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strconv"
)

// Download 下载文件
func Download(url string) (*bytes.Buffer, error) {
	// 下载文件
	body, err := Fetch(url)
	if err != nil {
		return nil, err
	}

	file := bytes.NewBuffer(nil)
	// 写入文件
	_, err = io.Copy(file, body)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Fetch 获取文件
func Fetch(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}

	body := response.Body

	return body, nil
}

// 判断参数是否为指针
func isPointer(v reflect.Type) bool {
	t := v
	return t.Kind() == reflect.Ptr
}

// 是否是切片
func isSlice(v reflect.Type) bool {
	t := v
	return t.Kind() == reflect.Slice
}

// 判断参数是否为结构体
func isStruct(v reflect.Type) bool {
	t := v
	return t != nil && t.Kind() == reflect.Struct
}

// 判断参数是否是指针类型的切片, 且切片元素必须是指针类型的结构体
func isSliceStructPointer(v reflect.Type) bool {
	t := v
	return t != nil &&
		t.Kind() == reflect.Ptr &&
		t.Elem().Kind() == reflect.Slice &&
		t.Elem().Elem().Kind() == reflect.Ptr &&
		t.Elem().Elem().Elem().Kind() == reflect.Struct
}

// toInt 转换为整数
func toInt(v string) (int, error) {
	return strconv.Atoi(v)
}

// toFloat 转换为浮点数
func toFloat(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}
