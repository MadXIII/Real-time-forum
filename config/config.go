package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server `yaml:"server"`
	DB     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type DB struct {
	Driver   string `yaml:"driver"`
	Store    string `yaml:"store"`
	Filename string `yaml:"filename"`
}

func Get(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	config := &Config{}
	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
