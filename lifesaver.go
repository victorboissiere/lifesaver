package main

import (
	"./installer"
	"fmt"
	"os"
)

func printUsage(config installer.Config) {
	fmt.Printf("Usage: %s [INSTALL_COMMAND]\nAvailables modes:\n", os.Args[0])
	for mode, installation := range config {
		fmt.Printf("  - %s : %v\n", mode, installation.Description)
	}

	os.Exit(1)
}

func main() {
	config := installer.GetConfig()

	if len(os.Args) != 2 {
		printUsage(config)
	}

	installCommand := os.Args[1]
	_, isValid := config[installCommand]
	if !isValid {
		printUsage(config)
	}

	installer.InstallConfig(config, installCommand)
	fmt.Println("\nDone! Thank you for using LifeSaver")
}
