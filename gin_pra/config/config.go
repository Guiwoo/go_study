package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

var cfg *Config

type Config struct {
	Host string `yaml:"Host"`
	Mode string `yaml:"Mode"`
}

func SetConfig(path string) *Config {
	once := sync.Once{}
	once.Do(func() {
		if cfg == nil {
			data, err := os.ReadFile(path)
			if err != nil {
				// change to real gin_log instance
				log.Panicf("fail to read file %+v path %+v", err, path)
			}
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				log.Panicf("fail to unmarshal yaml file %+v path %+v", err, path)
			}
		}
	})
	return cfg
}
