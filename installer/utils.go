package installer

import (
	"os/exec"
	"log"
	"os"
	"fmt"
	"path/filepath"
)

func getRepoFileURL(filename string) string {
	return fmt.Sprintf("%s/%s", BASE_CONFIG_ASSETS_URL, filename)
}

func execCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("Stdout: %sFailed with %s\n", out, err)
	}
}

func DownloadFile(url string, filename string) {
	execCommand(fmt.Sprintf("wget -O %s %s", filename, url))
}

func createPathIfNotExists(filename string) {
	err := os.MkdirAll(filepath.Dir(filename), 0744)
	if err != nil {
		log.Fatalf("Error creating path: %s\n", err)
	}
}
