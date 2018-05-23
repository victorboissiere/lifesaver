package installer

import (
	"fmt"
	"strings"
)

func importConfigFiles(configFiles []ConfigFile) {
	fmt.Println("[STEP][CONFIG_FILES]")

	for _, configFile := range configFiles {
		srcPath := getRepoFileURL(configFile.Src)
		dstPath := resolveTilde(configFile.Dst)
		fmt.Printf("[STEP][CONFIG_FILE] %s => %s\n", srcPath, dstPath)
		DownloadFile(srcPath, dstPath)
		setOwnership(dstPath)
	}
}

func getPostInstallHelp(message string) string {
	if len(message) != 0 {
		return fmt.Sprintf("\n\tHelp: %s\n", message)
	}

	return ""
}

func installCommands(commands []string) {
	for _, command := range commands {
		commandResolved := resolveTilde(command)
		fmt.Printf("[COMMAND] %s\n", commandResolved)
		execCommand(commandResolved)
	}
}

func installSteps(steps []InstallStep) {
	fmt.Println("[STEPS]")
	for _, step := range steps {
		fmt.Printf("[STEP] %s\n", step.Description)

		importConfigFiles(step.ConfigFiles)
		installCommands(step.Commands)
	}
}

func installPrograms(programs []string) {
	fmt.Println("[PROGRAMS]")
	for _, program := range programs {
		fmt.Printf("[PROGRAM] %s\n", strings.ToUpper(program))
		execCommand("apt install -y " + program)
	}
}

func install(installation Installation)  {
	installPrograms(installation.Programs)
	installSteps(installation.Steps)
}

func InstallConfig(config Config, installCommand string) {
	postInstallHelp := ""
	installation := config[installCommand]

	for _, dependency := range installation.Dependencies {
		fmt.Printf("====> Installing dependency '%s'\n", dependency)
		install(config[dependency])
		postInstallHelp += getPostInstallHelp(config[dependency].AfterHelp)
	}

	fmt.Printf("====> Installing '%s'\n", installCommand)
	install(installation)
	fmt.Printf("%s%s", postInstallHelp, getPostInstallHelp(installation.AfterHelp))
}

