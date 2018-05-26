package installer

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

const BaseConfigAssetsUrl = "https://raw.githubusercontent.com/victorboissiere/lifesaver/master"

type Config map[string]Installation
type Installation struct {
	Description  string
	Programs     []string
	Dependencies []string // optional
	Steps        []InstallStep
	AfterHelp    string `yaml:"afterHelp"`
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
	configFileURL := getRepoFileURL("config.yaml")

	fmt.Printf("[CONFIG] Repo is %s\n", getRepoFileURL(""))
	fmt.Printf("[CONFIG] Downloading config file %s\n", configFileURL)

	tmpConfigFile := "/tmp/lifesaver_config.yaml"
	DownloadFile(configFileURL, tmpConfigFile)

	yamlFile, err := ioutil.ReadFile(tmpConfigFile)
	if err != nil {
		log.Fatalln(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalln(err)
	}

	os.Remove(tmpConfigFile)

	return config
}
