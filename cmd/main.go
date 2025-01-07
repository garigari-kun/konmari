package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "konmari",
		Usage: "Organize your directory by file extensions",
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return fmt.Errorf("Please provide a directory path")
			}

			dirPath := c.Args().Get(0)

			return printFilesInDirectory(dirPath)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printFilesInDirectory(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("Directory does not exist: %s", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		fmt.Println(filepath.Join(dirPath, entry.Name()))
	}

	return nil
}
