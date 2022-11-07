package load

import (
	"gopkg.in/yaml.v2"
	"os"
)

// GetConf yml映射结构体
func GetConf[T any](filePath string, confStructPtr *T) error {

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if yaml.Unmarshal(content, confStructPtr) != nil {
		return err
	}

	return nil
}
