package installer

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Config map[string]Installation
type Installation struct {
	Description  string
	Programs     []string
	Dependencies []string // optional
	Steps []InstallStep
	AfterHelp string `yaml:"afterHelp"`
}
type InstallStep struct {
	Description string
	ConfigFiles []ConfigFile `yaml:"configFiles"`
	Commands    []string
}
type ConfigFile struct {
	Src string
	Dst string
}


func GetConfig() Config {
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
