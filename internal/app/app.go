// Package app is the main package for the application.
package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/allalt/internal/ai"
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
			Usage:   "the API key to use",
			EnvVars: []string{
				"ALLALT_KEY",
			},
		},
		&cli.StringFlag{
			Name:    "provider",
			Aliases: []string{"p"},
			Usage:   "the AI provider to use",
			Value:   ai.ProviderAnthropic,
			EnvVars: []string{
				"ALLALT_PROVIDER",
			},
		},
		&cli.StringFlag{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "the model to use when describing images",
			Value:   "claude-3-5-sonnet-20240620",
			EnvVars: []string{
				"ALLALT_MODEL",
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
		&cli.BoolFlag{
			Name:    "filename",
			Aliases: []string{"f"},
			Usage:   "whether to generate a filename for the image",
			EnvVars: []string{
				"ALLALT_FILENAME",
			},
		},
	}

	app.Action = DescribeAction

	if err := app.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		return 1
	}

	return 0
}
