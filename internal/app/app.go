// Package app is the main package for the application.
package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/allalt/internal/meta"
	"github.com/urfave/cli/v2"
)

// Run is the entry point for the application.
func Run(args []string) int {
	app := cli.NewApp()
	app.Name = meta.Name
	app.Version = meta.Version
	app.Usage = meta.Description
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "key",
			Aliases: []string{"k"},
			Usage:   "the OpenAI API key to use",
			EnvVars: []string{
				"ALLALT_KEY",
			},
		},
		&cli.StringFlag{
			Name:    "language",
			Aliases: []string{"l"},
			Usage:   "the language to use when describing images",
			Value:   "en_US",
			EnvVars: []string{
				"ALLALT_LANGUAGE",
			},
		},
		&cli.StringFlag{
			Name:    "context",
			Aliases: []string{"c"},
			Usage:   "the context around the image to use when describing images",
			EnvVars: []string{
				"ALLALT_CONTEXT",
			},
		},
		&cli.StringSliceFlag{
			Name:    "keyword",
			Aliases: []string{"K"},
			Usage:   "potential keywords relevant to the image",
		},
	}

	app.Action = DescribeAction

	if err := app.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		return 1
	}

	return 0
}
