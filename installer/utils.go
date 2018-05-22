package installer

import (
	"os/exec"
	"log"
	"os"
	"strconv"
	"os/user"
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
)

func getRepoFileURL(filename string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/victorboissiere/lifesaver/go/%s", filename)
}

func resolveTilde(command string) string {
	usr, err := user.LookupId(getUID())
	if err != nil {
		log.Fatal( err )
	}

	return strings.Replace(command, "~", usr.HomeDir, -1)
}

func execCommand(command string) {
	cmd := exec.Command(os.Getenv("SHELL"), "-c", command)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("Stdout: %sFailed with %s\n", out, err)
	}
}

func DownloadFile(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	contents, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		fmt.Printf("%s", errBody)
		os.Exit(1)
	}

	err = ioutil.WriteFile(resolveTilde(filename), contents, 0644)
	if err != nil {
		return err
	}

	return nil
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

