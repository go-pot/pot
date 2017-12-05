package run

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/wzshiming/pot"
)

var (
	appname      = ""
	cmd          *exec.Cmd
	state        sync.Mutex
	scheduleTime time.Time
	watchExts    = []string{".go"}
)

func init() {
	wd, _ := os.Getwd()
	appname = filepath.Base(wd)
}

// NewWatcher starts an fsnotify Watcher on the specified paths
func NewWatcher(paths []string, cb func()) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case e := <-watcher.Events:

				if !shouldWatchFileWithExtension(e.Name) {
					continue
				}

				pot.Println("Event fired: ", e)
				if cb != nil {
					cb()
				}
			loop:
				for {
					select {
					case <-watcher.Events:
					default:
						break loop
					}
				}
			case err := <-watcher.Errors:
				pot.Println("Watcher error: %s", err.Error()) // No need to exit here
			}
		}
	}()

	pot.Println("Initializing watcher...")

	m := map[string]int{}

	for _, path := range paths {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasPrefix(path, ".") {
				return nil
			}

			if !shouldWatchFileWithExtension(path) {
				return nil
			}

			dir := filepath.Dir(path)
			m[dir]++
			return nil
		})
	}

	for path, _ := range m {
		pot.Println("Watching: ", path)
		err = watcher.Add(path)
		if err != nil {
			pot.Println("Failed to watch directory: %s", err)
		}
	}

	<-done
	return nil
}

func GoGenerate() {
	cmdName := "go"
	args := []string{"generate", "./..."}
	bcmd := exec.Command(cmdName, args...)
	var stderr bytes.Buffer
	bcmd.Stderr = &stderr
	err := bcmd.Run()
	if err != nil {
		pot.Println("Failed to generate the application: %s", stderr.String())
		return
	}

	pot.Println("Generate Successfully!")
}

// AutoBuild builds the specified set of files
func GoBuild(cb func()) {
	state.Lock()
	defer state.Unlock()
	if cb != nil {
		cb()
	}

	cmdName := "go"
	appName := appname
	if runtime.GOOS == "windows" {
		appName += ".exe"
	}
	args := []string{"build"}
	args = append(args, "-o", appName)

	bcmd := exec.Command(cmdName, args...)
	var stderr bytes.Buffer
	bcmd.Stderr = &stderr
	err := bcmd.Run()
	if err != nil {
		pot.Println("Failed to build the application: %s", stderr.String())
		return
	}

	pot.Println("Built Successfully!")
	Restart(appname)
}

// Kill kills the running command process
func Kill() {
	defer func() {
		if e := recover(); e != nil {
			pot.Println("Kill recover: ", e)
		}
	}()
	if cmd != nil && cmd.Process != nil {
		err := cmd.Process.Kill()
		if err != nil {
			pot.Println("Error while killing cmd process: ", err)
		}
	}
}

// Restart kills the running command process and starts it again
func Restart(appname string) {
	pot.Println("Kill running process")
	Kill()
	go Start(appname)
}

// Start starts the command process
func Start(appname string) {
	pot.Println("Restarting ", appname)
	if !strings.Contains(appname, "./") {
		appname = "./" + appname
	}

	cmd = exec.Command(appname)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go cmd.Run()
	pot.Println(appname, " is running...")

}

// shouldWatchFileWithExtension returns true if the name of the file
// hash a suffix that should be watched.
func shouldWatchFileWithExtension(name string) bool {
	for _, s := range watchExts {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

// GetFileModTime returns unix timestamp of `os.File.ModTime` for the given path.
func GetFileModTime(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	f, err := os.Open(path)
	if err != nil {
		pot.Println("Failed to open file on '%s': %s", path, err)
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		pot.Println("Failed to get file stats: %s", err)
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
