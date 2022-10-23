package shell

import (
	"strings"
)

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func filterEmptyLines(lines []string) []string {
	var out []string = []string{}
	for _, line := range lines {
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}

func trimLines(lines []string) []string {
	var out []string = []string{}
	for _, line := range lines {
		out = append(out, strings.TrimSpace(line))
	}
	return out
}
