package installer

import (
	"os/exec"
	"log"
	"os"
	"strconv"
	"os/user"
	"strings"
)

func resolveTilde(command string) string {
	usr, err := user.LookupId(getUID())
	if err != nil {
		log.Fatal( err )
	}

	return strings.Replace(command, "~", usr.HomeDir, -1)
}

func execCommand(command string) {
	shellCommand := strings.Split(resolveTilde(command), " ")
	cmd := exec.Command(shellCommand[0], shellCommand[1:]...)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("Stdout: %sFailed with %s\n", out, err)
	}
}

func getEnvNumber(env string) int {
	envInt, err := strconv.Atoi(env)
	if err != nil {
		log.Fatalf("Error converting SUDO_GID: %s\n", err)
	}

	return envInt
}

func getUID() string {
	if uid, ok := os.LookupEnv("SUDO_UID"); ok {
		return uid
	}

	return strconv.Itoa(os.Getuid())
}

func getGID() string {
	if gid, ok := os.LookupEnv("SUDO_GID"); ok {
		return gid
	}

	return strconv.Itoa(os.Getgid())
}

func setOwnership(filename string) {
	err := os.Chown(resolveTilde(filename), getEnvNumber(getUID()), getEnvNumber(getGID()))
	if err != nil {
		log.Fatalf("Error setting permissions: %s\n", err)
	}
}

