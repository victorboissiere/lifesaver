package installer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"net/http"
	"io/ioutil"
	"strings"
)

var (
	homePath string
)

func init() {
	homeEnv, ok := os.LookupEnv("HOME"); if !ok {
		log.Fatal("Could not get HOME environment")
	}
	homePath = homeEnv
}

func getRepoFileURL(filename string) string {
	return fmt.Sprintf("%s/%s", BaseConfigAssetsUrl, filename)
}

func execCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("Stdout: %sFailed with %s\n", out, err)
	}
}

func createPathIfNotExists(filename string) {
	err := os.MkdirAll(filepath.Dir(filename), 0744)
	if err != nil {
		log.Fatalf("Error creating path: %s\n", err)
	}
}

func resolveTilde(command string) string {
	return strings.Replace(command, "~", homePath, -1)
}

func DownloadFile(url string, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal( fmt.Sprintf("Could not fetch remote file %s: %s", url, err ))
	}
	defer resp.Body.Close()

	contents, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		fmt.Printf("%s", errBody)
		os.Exit(1)
	}

	err = ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		log.Fatal( fmt.Sprintf("Could not write to file '%s': %s\n", filename, err ))
	}
}

