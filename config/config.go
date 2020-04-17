package config

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

type Conf struct {
	Service struct {
		Addr string `yaml:"addr"`
		Name string `yaml:"name"`
		Log  string `json:"log"`
	} `yaml:"server"`

	// todo:: may be array
	Mysql struct {
		Addr     string `yaml:"addr"`
		Port     int    `yaml:"port"`
		DbName   string `yaml:"db_name"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`

	Redis struct {
		Key  string `yaml:"key"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Auth string `yaml:"auth"`
		Db   int    `yaml:"db"`
	} `yaml:"redis"`

	Consul string `yaml:"consul"`

	Jaeger string `yaml:"jaeger"`

	Extra interface{} `yaml:"extra"`
}

var C Conf

func Init(path string, extra interface{}) {
	if _, err := os.Stat(path); err != nil {
		panic("find yaml fail: " + err.Error())
	}

	bts, err := ioutil.ReadFile(path)
	if err != nil {
		panic("read yaml fail: " + err.Error())
	}
	C.Extra = extra
	err = yaml.Unmarshal(bts, &C)
	if err != nil {
		panic("yaml unmarshal fail: " + err.Error())
	}
}
