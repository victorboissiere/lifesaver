package main

import "fmt"

func importConfigFiles(configFiles []ConfigFile) {
	fmt.Println("[STEP][CONFIG_FILES]")
	for _, configFile := range configFiles {
		fmt.Printf("[STEP[CONFIG_FILE] %s => %s\n", configFile.Src, configFile.Dst)
		execCommand(fmt.Sprintf("wget -O %s https://raw.githubusercontent.com/victorboissiere/lifesaver/master/%s", configFile.Dst, configFile.Src))
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


