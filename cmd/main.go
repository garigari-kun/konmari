package main

import (
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
				return cli.Exit("Please provide a directory path", 1)
			}
			dirPath := c.Args().Get(0)
			return organizer.OrganizeFiles(dirPath)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
