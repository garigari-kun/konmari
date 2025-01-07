package main

import (
	"fmt"
	"log"
	"os"

	"github.com/garigari-kun/konmari/internal/organizer"
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
			return organizer.OrganizeFiles(dirPath)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
