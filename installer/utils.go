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
	"path/filepath"
	"syscall"
)

func getRepoFileURL(filename string) string {
	return fmt.Sprintf("%s/%s", BASE_CONFIG_ASSETS_URL, filename)
}

func resolveTilde(command string) string {
	usr, err := user.LookupId(strconv.Itoa(getUID()))
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

	err = ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getEnvNumber(env string) int {
	envInt, err := strconv.Atoi(env)
	if err != nil {
		log.Fatalf("Error converting %s: %s\n", env, err)
	}

	return envInt
}

func getUID() int {
	if uid, ok := os.LookupEnv("SUDO_UID"); ok {
		return getEnvNumber(uid)
	}

	return os.Getuid()
}

func getGID() int {
	if gid, ok := os.LookupEnv("SUDO_GID"); ok {
		return getEnvNumber(gid)
	}

	return os.Getgid()
}

func setOwnership(filename string) {
	err := os.Chown(filename, getUID(), getGID())
	if err != nil {
		log.Fatalf("Error setting permissions: %s\n", err)
	}
}


func createPathIfNotExists(filename string) {
	err := mkdirAll(filepath.Dir(filename), 0744)
	if err != nil {
		log.Fatalf("Error creating path: %s\n", err)
	}
}

// From original lib because need to chown as well as process usually run as sudo
func mkdirAll(path string, perm os.FileMode) error {
	dir, err := os.Stat(path)
	if err == nil {
		if dir.IsDir() {
			return nil
		}
		return &os.PathError{"mkdir", path, syscall.ENOTDIR}
	}

	i := len(path)
	for i > 0 && os.IsPathSeparator(path[i-1]) {
		i--
	}

	j := i
	for j > 0 && !os.IsPathSeparator(path[j-1]) {
		j--
	}

	if j > 1 {
		err = mkdirAll(path[0:j-1], perm)
		if err != nil {
			return err
		}
	}

	err = os.Mkdir(path, perm)
	if err != nil {
		dir, err1 := os.Lstat(path)
		if err1 == nil && dir.IsDir() {
			return nil
		}
		return err
	}
	// Added function not in lib
	setOwnership(path)

	return nil
}

