package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	configPathDefault = "config.yml"
)

// Config is global config file type
type Config struct {
	Addr   string `yaml:"addr"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
}

var (
	// Cfg is global config in disk
	Cfg Config
)

func init() {
	loadConfigure()
}

func loadConfigure() {
	buf, err := ioutil.ReadFile(configPathDefault)
	if err != nil {
		panic("config file path error")
	}

	err = yaml.Unmarshal(buf, &Cfg)
	if err != nil {
		panic("yaml config file parse error")
	}

	fmt.Println(Cfg)
}
