package shell

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	blacklist "github.com/lordsalmon/monitormanagement/blacklist"
	database "github.com/lordsalmon/monitormanagement/database"
	log "github.com/sirupsen/logrus"
)

func GetAllWindows() []database.Window {
	var windows []database.Window = []database.Window{}
	if runtime.GOOS == "linux" {
		windows = getLinuxWindows()
	} else if runtime.GOOS == "darwin" {
		windows = getMacWindows()
	}
	return windows
}

func getLinuxWindows() []database.Window {
	log.Info("Getting windows...")
	cmd := exec.Command("bash", "-c", "xwininfo -root -tree | grep 0x")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	var lines []string = strings.Split(string(stdout), "\n")
	lines = trimLines(lines)
	lines = filterLinesByBlacklist(lines)
	lines = filterEmptyLines(lines)
	var windows []database.Window = []database.Window{}
	for _, line := range lines {
		var window database.Window = database.Window{}
		windowId, err := strconv.ParseInt(strings.Split(line, " ")[0], 0, 64)
		if err != nil {
			log.Error("Error parsing window id:", err)
			os.Exit(1)
		}
		// boilerplate line: 0x8a00008 "Discord": ("Discord" "Discord")  16x16+0+0  +0+0
		window.WindowId = int(windowId)
		program := strings.Split(line, "(\"")[1]
		program = strings.Split(program, "\")")[0]
		program = strings.Split(program, "\"")[2]
		window.Program = program
		window = InjectWindowInformation(window)
		windows = append(windows, window)
	}
	return windows
}

func getMacWindows() []database.Window {
	return []database.Window{}
}

func filterLinesByBlacklist(lines []string) []string {
	var out []string = []string{}
	for _, line := range lines {
		if !isBlacklisted(line) {
			out = append(out, line)
		}
	}
	return out
}

func isBlacklisted(line string) bool {
	for _, blacklistEntry := range blacklist.Blacklist {
		if strings.Contains(line, blacklistEntry) {
			return true
		}
	}
	return false
}
