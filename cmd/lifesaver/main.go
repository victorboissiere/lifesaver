package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Config map[string]struct {
	Description string
	Programs []string
	Dependencies []string // optional
	Steps []struct {
		Description string
		Commands interface{}
	}
}

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

func printUsage(config Config) {
	fmt.Printf("Usage: %s [INSTALL_COMMAND]\nAvailables modes:\n", os.Args[0])
	for mode, installation := range config {
		fmt.Printf("  - %s : %v\n", mode, installation.Description)
	}

	os.Exit(1)
}

func main()  {
	config := getConfig()

	if len(os.Args) != 2 {
		printUsage(config)
	}
	installCommand := os.Args[1]

	fmt.Printf("Install mode: %s", installCommand)
}