package main

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

func getConfig() Config {
	yamlFile, err := ioutil.ReadFile("config2.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalln(err)
	}

	return config
}
