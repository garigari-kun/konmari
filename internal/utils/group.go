package utils

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GroupFilesByExtension(entries []fs.DirEntry) map[string][]string {
	fileGroups := make(map[string][]string)

	for _, entry := range entries {
		if entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext == "" {
			ext = "no_extension"
		} else {
			ext = ext[1:]
		}

		fileGroups[ext] = append(fileGroups[ext], entry.Name())
	}

	return fileGroups
}
