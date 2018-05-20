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
	ConfigFiles []ConfigFile `yaml:"configFiles"`
	Commands    []string
}
type ConfigFile struct {
	Src string
	Dst string
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

func getUsername() string {
	if user, ok := os.LookupEnv("SUDO_USER"); ok {
		return user
	}

	return os.Getenv("USER")
}

func setOwnership(filename string) {
	username := getUsername()
	execCommand(fmt.Sprintf("chown -R %s:%s %s", username, username, filename))
}

func importConfigFiles(configFiles []ConfigFile) {
	fmt.Println("[STEP][CONFIG_FILES]")
	for _, configFile := range configFiles {
		fmt.Printf("[STEP[CONFIG_FILE] %s => %s\n", configFile.Src, configFile.Dst)
		execCommand(fmt.Sprintf("wget https://raw.githubusercontent.com/victorboissiere/lifesaver/master/%s -O %s", configFile.Src, configFile.Dst))
		setOwnership(configFile.Dst)
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