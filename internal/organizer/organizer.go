package organizer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func OrganizeFiles(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("Directory does not exist: %s", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

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

	for ext, files := range fileGroups {
		subDir := filepath.Join(dirPath, ext)
		if err := os.MkdirAll(subDir, os.ModePerm); err != nil {
			return fmt.Errorf("Failed to create directory %s: %v", subDir, err)
		}
		fmt.Printf("Created directory: %s\n", subDir)

		for _, file := range files {
			oldPath := filepath.Join(dirPath, file)
			newPath := filepath.Join(subDir, file)
			if err := os.Rename(oldPath, newPath); err != nil {
				return fmt.Errorf("Failed to move file %s to %s: %v", oldPath, newPath, err)
			}
			fmt.Printf("Moved: %s -> %s\n", oldPath, newPath)
		}
	}
	return nil
}
