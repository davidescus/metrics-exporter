package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Port int `yaml:"port"`
	CollectInterval int64 `yaml:"collectInterval"`
	TimeToLive uint64 `yaml:"timeToLive"`
	StoragePath string `yaml:"storagePath"`
}

// NewConf...
func NewConf(path string) *Conf {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Configuration file could not be read from: " + path)
	}

	c := &Conf{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}

	return c
}
