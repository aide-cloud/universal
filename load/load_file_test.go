package load

import "testing"

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
	conf, err := GetConf[Conf]("test.yml")
	if err != nil {
		t.Error(err)
	}
	t.Log(conf)
}
