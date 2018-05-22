package main

import (
	"os"
	"fmt"
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


func printUsage(config Config) {
	fmt.Printf("Usage: %s [INSTALL_COMMAND]\nAvailables modes:\n", os.Args[0])
	for mode, installation := range config {
		fmt.Printf("  - %s : %v\n", mode, installation.Description)
	}

	os.Exit(1)
}


func installPrograms(programs []string) {
	fmt.Println("[PROGRAMS]")
	for _, program := range programs {
		fmt.Printf("[PROGRAM] %s\n", strings.ToUpper(program))
		execCommand("apt install -y " + program)
	}
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
	fmt.Println("\nDone! Thank you for using LifeSaver")
}