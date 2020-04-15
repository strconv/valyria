package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

var Conf struct {
	Service struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	}
	// todo:: may be array
	Mysql struct {
		Address  string `yaml:"address"`
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

func InitConf(path string, extra interface{}) {
	if _, err := os.Stat(path); err != nil {
		panic("find yaml fail: " + err.Error())
	}

	bts, err := ioutil.ReadFile(path)
	if err != nil {
		panic("read yaml fail: " + err.Error())
	}
	Conf.Extra = extra
	err = yaml.Unmarshal(bts, &Conf)
	if err != nil {
		log.Panicf("yaml unmarshal fail|err:%s|row:%s|", err, bts)
		return
	}
}
