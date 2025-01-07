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
		if err := utils.MoveFiles(dirPath, subDir, files); err != nil {
			return err
		}
	}
	return nil
}
