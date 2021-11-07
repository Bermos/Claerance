package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

const configPath = "conf/conf.yml"

var Cfg Config

func Load() {
	absPath, _ := filepath.Abs(configPath)
	file, err := os.Open(absPath)
	if err != nil {
		log.Println("INFO - No config file found, using default values")
		log.Fatal(err)
		return
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&Cfg); err != nil {
		log.Fatal("ERROR - Malformed config.yml", err)
	}
}
