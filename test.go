package main

import (
	. "./utils"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func test() {
	yamlfile, err := ioutil.ReadFile("./config.yaml")
	CheckErr(err)

	var config Config
	err = yaml.Unmarshal(yamlfile, &config)
	CheckErr(err)

	fmt.Println(config.Dev.Port)
}
