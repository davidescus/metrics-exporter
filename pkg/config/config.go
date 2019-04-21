package config

import (
	"errors"
	"github.com/crgimenes/goconfig"
	_ "github.com/crgimenes/goconfig/yaml"
)

var defaultConfigFile = "./configs/exporter.yaml"

type Config struct {
	Port int
	//scraping interval in seconds
	ScrapeInterval int `yaml:"scrapeInterval"`
	// maximum duration for a running collection process
	// kill after exceed time
	MaxProcessCollectDuration int `yaml:"maxProcessCollectDuration"`
	// maximum amount of data to store
	// will truncate old values
	// value have to be in bytes
	MaxAmountOfDataToStore int `yaml:"maxAmountOfDataToStore"`
}

// NewConfig instantiate general configuration
func NewConfig() (*Config, error) {
	// TODO get configuration from many places (local, environment, etc, flags)
	configFile := defaultConfigFile

	config := Config{}
	goconfig.File = configFile

	err := goconfig.Parse(&config)
	if err != nil {
		return nil, err
	}

	// validate configuration
	if config.ScrapeInterval <= config.MaxProcessCollectDuration {
		return nil, errors.New("MaxProcessCollectDuration have to be less than ScrapeInterval")
	}

	return &config, nil
}
