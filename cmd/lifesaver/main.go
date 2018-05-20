package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"os/exec"
	"strings"
)

type Config map[string]Installation
type Installation struct {
	Description  string
	Programs     []string
	Dependencies []string // optional
	Steps []struct {
		Description string
		Commands    interface{}
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

func installPrograms(programs []string) {
	for _, program := range programs {
		fmt.Printf("[PROGRAM] %s\n", strings.ToUpper(program))
		cmd := exec.Command("apt", "install", "-y", program)
		if out, err := cmd.CombinedOutput(); err != nil {
			log.Fatalf("Stdout: %sFailed with %s\n", out, err)
		}
	}
}

func install(installation Installation) {
	fmt.Println("[PROGRAMS] Installing")
	installPrograms(installation.Programs)
}

func main()  {
	config := getConfig()

	if len(os.Args) != 2 {
		printUsage(config)
	}

	installCommand := os.Args[1]
	installation, isValid := config[installCommand]
	if !isValid {
		printUsage(config)
	}

	fmt.Printf("====> Installing '%s'\n", installCommand)
	install(installation)
}