package load

import (
	"encoding/json"
	"encoding/xml"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

// GetConf yml映射结构体
func GetConf(filePath string, confStructPtr any) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 文件类型
	fileTypeStr := path.Ext(filePath)

	switch fileTypeStr {
	case ".yml", ".yaml":
		return yaml.Unmarshal(content, confStructPtr)
	case ".json":
		return json.Unmarshal(content, confStructPtr)
	case ".toml":
		return toml.Unmarshal(content, confStructPtr)
	case ".xml":
		return xml.Unmarshal(content, confStructPtr)
	default:
		panic("file type error")
	}
}
