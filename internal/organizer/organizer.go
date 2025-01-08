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

func PreviewOrganizeFiles(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("Directory does not exist: %s", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	fileGroups := utils.GroupFilesByExtension(entries)

	fmt.Println("Preview of file organization:")
	for ext, files := range fileGroups {
		subDir := filepath.Base(filepath.Join(dirPath, ext))
		fmt.Printf("├── %s/\n", subDir)
		for i, file := range files {
			prefix := "|"
			if i == len(files)-1 {
				prefix = "└"
			}
			fmt.Printf("%s   %s\n", prefix, file)
		}
	}

	return nil
}
