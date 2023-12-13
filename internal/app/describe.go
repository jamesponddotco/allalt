package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"git.sr.ht/~jamesponddotco/allalt/internal/openai"
	"git.sr.ht/~jamesponddotco/allalt/internal/xbase64"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyInput is returned when no image is provided by the user.
	ErrEmptyInput xerrors.Error = "missing image to describe"

	// ErrEmptyKey is returned when no API key for OpenAI is provided by the user.
	ErrEmptyKey xerrors.Error = "missing OpenAI API key"
)

// DescribeAction is the main action for the application.
func DescribeAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return ErrEmptyInput
	}

	var (
		key      = ctx.String("key")
		language = ctx.String("language")
		keywords = ctx.StringSlice("keyword")
		image    = ctx.Args().Get(0)
	)

	if key == "" {
		return ErrEmptyKey
	}

	data, err := os.ReadFile(image)
	if err != nil {
		return fmt.Errorf("failed to read image: %w", err)
	}

	var (
		base64Image = xbase64.EncodeImageToDataURL(data)
		client      = openai.NewClient(key)
		req         = openai.NewRequest(language, base64Image, keywords)
	)

	resp, err := client.Do(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}

	for {
		text, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Fprintf(ctx.App.Writer, "\n")

			break
		}

		if err != nil {
			return fmt.Errorf("failed to get response: %w", err)
		}

		fmt.Fprintf(ctx.App.Writer, "%s", text.Choices[0].Delta.Content)
	}

	return nil
}
