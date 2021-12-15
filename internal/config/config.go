package config

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Api struct {
		Host string `yaml:"host"`
	} `yaml:"api"`

	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string
	} `yaml:"mysql"`
}

func GetConfig(pathConfig string) (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(pathConfig)
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

	config.MySQL.Password = os.Getenv("MYSQL_PASSWORD")

	return config, nil
}
