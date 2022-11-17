package load

import (
	"fmt"
	"testing"
)

type Conf struct {
	Server struct {
		Addr string `yaml:"addr" json:"addr"`
		Port int    `yaml:"port" json:"port"`
	}

	Mysql struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		User   string `yaml:"user"`
		Pass   string `yaml:"password" json:"password"`
		DbName string `yaml:"database" json:"database"`
	}

	Redis struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		Pass   string `yaml:"password" json:"password"`
		DbName string `yaml:"database" json:"database"`
	}
}

func TestGetConfYaml(t *testing.T) {
	var conf Conf
	err := GetConf("test.yml", &conf)
	if err != nil {
		t.Error(err)
	}
	t.Log(conf)
	fmt.Printf("%p %+v", &conf, conf)
}

func TestGetConfJSON(t *testing.T) {
	var conf Conf
	err := GetConf("test.json", &conf)
	if err != nil {
		t.Error(err)
	}
	t.Log(conf)
	fmt.Printf("%p %+v", &conf, conf)
}
