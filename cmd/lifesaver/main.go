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
	Steps []InstallStep
}
type InstallStep struct {
	Description string
	ConfigFiles []ConfigFile
	Commands    []string
}
type ConfigFile struct {
	src string
	dst string
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

func execCommand(command string) {
	shellCommand := strings.Split(command, " ")
	cmd := exec.Command(shellCommand[0], shellCommand[1:]...)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("Stdout: %sFailed with %s\n", out, err)
	}
}

func installPrograms(programs []string) {
	fmt.Println("[PROGRAMS]")
	for _, program := range programs {
		fmt.Printf("[PROGRAM] %s\n", strings.ToUpper(program))
		execCommand("apt install -y " + program)
	}
}

func importConfigFiles(configFiles []ConfigFile) {
	fmt.Println("[STEP][CONFIG_FILES]")
	for _, configFile := range configFiles {
		fmt.Printf("[STEP[CONFIG_FILE] %s => %s\n", configFile.src, configFile.dst)
		execCommand(fmt.Sprintf("wget -O - https://raw.githubusercontent.com/victorboissiere/lifesaver/master/softwares/%s > %s", configFile.src, configFile.dst))
	}
}

func installSteps(steps []InstallStep) {
	fmt.Println("[STEPS]")
	for _, step := range steps {
		fmt.Printf("[STEP] %s\n", step.Description)
		importConfigFiles(step.ConfigFiles)
		for _, command := range step.Commands {
			execCommand(command)
		}
	}
}

func install(installation Installation) {
	installPrograms(installation.Programs)
	installSteps(installation.Steps)
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
	fmt.Println("Done! Thank you for using LifeSaver")
}