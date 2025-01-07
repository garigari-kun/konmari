package utils

import (
	"fmt"
	"os"
)

func CreateDirectory(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %v", path, err)
	}
	fmt.Printf("Created directory: %s\n", path)
	return nil
}
