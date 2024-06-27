package app

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const defaultConfigPath = "./config/config.yml"

type Config struct {
	Name    string `yaml:"name" default:"card_detector"`
	Version string `yaml:"version" default:"1.0.0"`

	Port int `yaml:"port" default:"8080"`

	TempFolder    string `yaml:"tmp_folder" default:"./tmp"`
	StorageFolder string `yaml:"storage_folder" default:"./storage"`
}

func NewConfigFromYml() (*Config, error) {
	configFilePath := os.Getenv("CONFIG_FILE")
	if configFilePath == "" {
		configFilePath = defaultConfigPath
	}
	var config Config
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Printf("Error reading YAML file: %s\n", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Error parsing YAML file: %s\n", err)
		return nil, err
	}

	return &config, nil
}
