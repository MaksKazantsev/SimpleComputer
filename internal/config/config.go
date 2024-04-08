package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port int `yaml:"port"`
}

func MustLoad() *Config {
	path := fetchPath()
	if path == "" {
		panic("config file path is required")
	}

	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read config file: " + err.Error())
	}

	var cfg Config
	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to parse config file: " + err.Error())
	}

	return &cfg
}

func fetchPath() string {
	var path string
	flag.StringVar(&path, "c", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = "./config.yaml"
	}

	return path
}
