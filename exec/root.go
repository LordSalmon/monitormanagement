package monitormanagement

import (
	"fmt"
	database "monitormanagement/database"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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
	cmd := exec.Command("xwininfo", "-root", "-tree", "|", "grep 0x")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	var lines []string = strings.Split(string(stdout), "\n")
	// remove first three lines as they are boilerplate
	lines = remove(lines, 0)
	lines = remove(lines, 0)
	lines = remove(lines, 0)
	for index, line := range lines {
		lines[index] = strings.ReplaceAll(lines[index], "'", "")
		lines[index] = strings.TrimSpace(line)
	}

	var windows []database.Window = []database.Window{}
	for _, line := range lines {
		var window database.Window = database.Window{}
		windowId, err := strconv.Atoi(strings.Split(line, " ")[0])
		if err != nil {
			fmt.Println("Error parsing window id:", err)
			os.Exit(1)
		}
		window.WindowId = windowId
		windows = append(windows, window)
	}
	return windows
}

func getMacWindows() []database.Window {
	return []database.Window{}
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
