package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
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
	setDefaultLogging()
	absPath, _ := filepath.Abs(configPath)
	file, err := os.Open(absPath)
	if err != nil {
		log.Infoln("No config file found, using default values")
		log.Fatalln("Default config not implemented. Abort")
		return
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&Cfg); err != nil {
		log.Fatalln("Malformed config.yml", err)
	}
}

func setDefaultLogging() {
	log.SetFormatter(&log.TextFormatter{})
}
