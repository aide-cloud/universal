package conf

import (
	"github.com/aide-cloud/universal/load"
	"sync"
)

const (
	DefaultPath = "./configs/config.yml"
)

type (
	Config struct {
		Server Server `yaml:"server"`
	}

	Server struct {
		Mode     string `yaml:"mode"`
		HttpAddr string `yaml:"http_addr"`
	}
)

var (
	config *Config
	once   sync.Once
)

// LoadConfig sets the config pointer.
func LoadConfig(filePath string) {
	if filePath == "" {
		filePath = DefaultPath
	}

	if config == nil {
		once.Do(func() {
			config = &Config{}
			err := load.GetConf[Config](filePath, config)
			if err != nil {
				panic(err)
			}
		})
	}
}

// GetConfig returns the config pointer.
func GetConfig() Config {
	if config == nil {
		LoadConfig(DefaultPath)
	}
	return *config
}
