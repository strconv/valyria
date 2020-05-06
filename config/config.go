package config

import (
	"io/ioutil"
	"os"
	"time"

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
		DBName   string `yaml:"dbName"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`

	Redis struct {
		Host        string `yaml:"host"`
		Auth        string `yaml:"auth"`
		DB          int    `yaml:"db"`
		MaxIdle     int    `yaml:"maxIdle"`
		IdleTimeout int32  `yaml:"idleTimeout"`
		MaxActive   int    `yaml:"maxActive"`
	} `yaml:"redis"`

	JWT struct {
		Secret  string        `yaml:"secret"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"jwt"`

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
