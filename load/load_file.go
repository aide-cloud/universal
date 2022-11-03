package load

import (
	"gopkg.in/yaml.v2"
	"os"
)

// GetConf yml映射结构体
func GetConf[T any](filePath string) (*T, error) {
	var confStructPtr T
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if yaml.Unmarshal(content, &confStructPtr) != nil {
		return nil, err
	}

	return &confStructPtr, nil
}
