package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Config configuration for http
type Config struct {
	Dev struct {
		Port string `yaml:"port"`
	}
	Product struct {
		Port string `yaml:"port"`
	}
}

// CheckErr check error
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// GetConfig get configuration
func (c *Config) GetConfig() *Config {
	yamlfile, err := ioutil.ReadFile("./config.yaml")
	CheckErr(err)

	err = yaml.Unmarshal(yamlfile, c)
	CheckErr(err)

	return c
}
