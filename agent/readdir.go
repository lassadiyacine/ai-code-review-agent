package agent

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadDir(dir string) (string, error) {
	var content strings.Builder

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			content.WriteString("// fichier: " + path + "\n")
			content.Write(data)
			content.WriteString("\n\n")
		}
		return nil
	})

	return content.String(), err
}
