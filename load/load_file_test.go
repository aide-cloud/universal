package load

import (
	"fmt"
	"testing"
)

type Conf struct {
	Server struct {
		Addr string `yaml:"addr"`
		Port int    `yaml:"port"`
	}

	Mysql struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		User   string `yaml:"user"`
		Pass   string `yaml:"password"`
		DbName string `yaml:"database"`
	}

	Redis struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		Pass   string `yaml:"password"`
		DbName string `yaml:"database"`
	}
}

func TestGetConf(t *testing.T) {
	var conf Conf
	err := GetConf("test.yml", &conf)
	if err != nil {
		t.Error(err)
	}
	t.Log(conf)
	fmt.Printf("%p %+v", &conf, conf)
}
