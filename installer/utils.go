package installer

import (
	"strings"
	"os/exec"
	"log"
	"os"
	"fmt"
)

func execCommand(command string) {
	shellCommand := strings.Split(command, " ")
	cmd := exec.Command(shellCommand[0], shellCommand[1:]...)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("Stdout: %sFailed with %s\n", out, err)
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

