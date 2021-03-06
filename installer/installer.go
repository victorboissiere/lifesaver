package installer

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func importConfigFiles(configFiles []ConfigFile) {
	for _, configFile := range configFiles {
		srcPath := getRepoFileURL(configFile.Src)
		dstPath := resolveTilde(configFile.Dst)
		fmt.Printf("[STEP][CONFIG_FILE] %s => %s\n", configFile.Src, dstPath)
		createPathIfNotExists(dstPath)
		DownloadFile(srcPath, dstPath)
	}
}

func getPostInstallHelp(message string) string {
	if len(message) != 0 {
		return fmt.Sprintf("\n\tHelp: %s", message)
	}

	return ""
}

func installCommands(commands []string) {
	for _, command := range commands {
		fmt.Printf("[STEP][COMMAND] %s\n", command)
		execCommand(command)
	}
}

func installSteps(steps []InstallStep) {
	for _, step := range steps {
		fmt.Printf("[STEP] %s\n", step.Description)

		importConfigFiles(step.ConfigFiles)
		installCommands(step.Commands)
	}
}

func installPrograms(programs []string) {
	for _, program := range programs {
		fmt.Printf("[PROGRAM] %s\n", strings.ToUpper(program))

		cmd := exec.Command("sudo", "apt", "install", "-y", program)
		if out, err := cmd.CombinedOutput(); err != nil {
			log.Fatalf("Stdout: %sFailed with %s\n", out, err)
		}
	}
}

func install(installation Installation) {
	installPrograms(installation.Programs)
	installSteps(installation.Steps)
}

func InstallConfig(config Config, installCommand string) {
	postInstallHelp := ""
	installation := config[installCommand]

	for _, dependency := range installation.Dependencies {
		fmt.Printf("\n====> Installing dependency '%s'\n\n", dependency)
		install(config[dependency])
		postInstallHelp += getPostInstallHelp(config[dependency].AfterHelp)
	}

	fmt.Printf("\n====> Installing '%s'\n\n", installCommand)
	install(installation)
	fmt.Printf("%s%s", postInstallHelp, getPostInstallHelp(installation.AfterHelp))
}
