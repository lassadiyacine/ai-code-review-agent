package agent

import (
	"strings"
)

var excludedFiles = []string{
	"go.mod",
	"go.sum",
	".gitignore",
	"README.md",
	"main.go",
	"gemini.go",
	"parser.go",
	"modes.go",
	"interactive.go",
	"filter.go",
}

func FilterDiff(diff string) string {
	blocks := strings.Split(diff, "\ndiff --git")
	var filtered []string

	for _, block := range blocks {
		if block == "" {
			continue
		}

		length := 100
		if len(block) < length {
			length = len(block)
		}

		excluded := false
		for _, file := range excludedFiles {
			if strings.Contains(block, file) {
				excluded = true
				break
			}
		}

		if !excluded {
			filtered = append(filtered, "diff --git"+block)
		}
	}

	return strings.Join(filtered, "")
}
