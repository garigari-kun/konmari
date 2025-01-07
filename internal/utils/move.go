package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func MoveFiles(srcDir, destDir string, files []string) error {
	for _, file := range files {
		oldPath := filepath.Join(srcDir, file)
		newPath := filepath.Join(destDir, file)

		if err := os.Rename(oldPath, newPath); err != nil {
			return fmt.Errorf("Failed to move file %s to %s: %v", oldPath, newPath, err)
		}
		fmt.Printf("Moved: %s -> %s\n", oldPath, newPath)
	}
	return nil
}
