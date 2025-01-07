package organizer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/garigari-kun/konmari/internal/utils"
)

func OrganizeFiles(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("Directory does not exist: %s", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	fileGroups := utils.GroupFilesByExtension(entries)

	for ext, files := range fileGroups {
		subDir := filepath.Join(dirPath, ext)
		if err := utils.CreateDirectory(subDir); err != nil {
			return err
		}

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
