package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Api struct {
		Host string `yaml:"host"`
	} `yaml:"api"`
}

func GetConfig() (Config, error) {
	file, err := os.Open("config/config.yml")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	yamlDecoder := yaml.NewDecoder(file)
	err = yamlDecoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
