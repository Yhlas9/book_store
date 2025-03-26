package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	Psql Psql `yaml:"psql"`
}

type Psql struct {
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
}						

func ParseConfig() AppConfig {
	data, err := os.ReadFile("config/prod.yaml")
	if err != nil {
		panic(err)
	}

	var config AppConfig

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}
