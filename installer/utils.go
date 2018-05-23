package installer

import (
	"os/exec"
	"log"
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os/user"
	"strings"
)

func resolveTilde(command string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}

	return strings.Replace(command, "~", usr.HomeDir, -1)
}

func getRepoFileURL(filename string) string {
	return fmt.Sprintf("%s/%s", BASE_CONFIG_ASSETS_URL, filename)
}

func execCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
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

func createPathIfNotExists(filename string) {
	err := os.MkdirAll(filepath.Dir(filename), 0744)
	if err != nil {
		log.Fatalf("Error creating path: %s\n", err)
	}
}
