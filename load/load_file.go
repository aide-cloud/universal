package load

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

const (
	YAML = "yml"
	YML  = "yml"
	JSON = "json"
	INI  = "ini"
	CONF = "conf"
)

// Unmarshal 加载配置文件
func Unmarshal(filePath string, object interface{}) {
	// read yml file
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// 获取文件类型
	fileSuffix := path.Ext(filePath)

	fileType := YML
	switch fileSuffix {
	case ".yml":
		fileType = YML
	case ".yaml":
		fileType = YAML
	case ".json":
		fileType = JSON
	case ".ini":
		fileType = INI
	case ".conf":
		fileType = CONF
	default:
		// 文件类型提示
		fmt.Println("Only YML, YAML, JSON, INI, and conf files are allowed.")
		panic(fmt.Errorf("unsupported file suffix: %s", fileSuffix))
	}

	// decode yml file
	decoder := viper.New()
	decoder.SetConfigType(fileType)
	err = decoder.ReadConfig(file)
	if err != nil {
		panic(err)
	}

	// decode yml file
	err = decoder.Unmarshal(object)
	if err != nil {
		panic(err)
	}
}
