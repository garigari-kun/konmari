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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "preview",
				Usage: "Show what changes would be made without moving files",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.Exit("Please provide a directory path", 1)
			}
			dirPath := c.Args().Get(0)
			if c.Bool("preview") {
				return organizer.PreviewOrganizeFiles(dirPath)
			}
			return organizer.OrganizeFiles(dirPath)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
