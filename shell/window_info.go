package shell

import (
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	database "github.com/lordsalmon/monitormanagement/database"
)

type InfoLine struct {
	key   string
	value string
}

func InjectWindowInformation(window database.Window) database.Window {
	if runtime.GOOS == "linux" {
		return getLinuxWindowInformation(window)
	} else if runtime.GOOS == "darwin" {
		return getMacWindowInformation(window)
	} else {
		window.Title = "Dont use windows"
		return window
	}
}

func getLinuxWindowInformation(window database.Window) database.Window {
	cmd := exec.Command("bash", "-c", "xwininfo -id "+strconv.Itoa(window.WindowId))
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	var lines []string = strings.Split(string(stdout), "\n")
	lines = filterEmptyLines(lines)
	titleLine := lines[0]
	title := strings.Split(titleLine, "\"")[1]
	lines = remove(lines, 0)
	lines = trimLines(lines)
	// remove line '-geometry intxint+int+int'
	newLines := []string{}
	for _, line := range lines {
		if !strings.Contains(line, "-geometry") {
			newLines = append(newLines, line)
		}
	}
	lines = newLines
	infos := make(map[string]string)
	for _, line := range lines {
		split := strings.Split(line, ":")
		key := split[0]
		key = strings.ToLower(key)
		key = strings.ReplaceAll(key, " ", "_")
		key = strings.ReplaceAll(key, "-", "_")
		value := split[1]
		value = strings.TrimSpace(value)
		infos[key] = value
	}
	depth, err := strconv.Atoi(infos["depth"])
	height, err := strconv.Atoi(infos["height"])
	width, err := strconv.Atoi(infos["width"])
	x, err := strconv.Atoi(infos["absolute_upper_left_x"])
	y, err := strconv.Atoi(infos["absolute_upper_left_y"])
	window.Title = title
	window.Depth = depth
	window.Height = height
	window.Width = width
	window.Y = y
	window.X = x

	return window
}

func getMacWindowInformation(window database.Window) database.Window {
	return window
}
