package app

import (
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~jamesponddotco/allalt/internal/ai"
	"git.sr.ht/~jamesponddotco/allalt/internal/ai/anthropic"
	"git.sr.ht/~jamesponddotco/allalt/internal/ai/openai"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyInput is returned when no image is provided by the user.
	ErrEmptyInput xerrors.Error = "missing image to describe"

	// ErrEmptyKey is returned when no API key for OpenAI is provided by the user.
	ErrEmptyKey xerrors.Error = "missing OpenAI API key"

	// ErrInvalidProvider is returned when an invalid provider is provided by the user.
	ErrInvalidProvider xerrors.Error = "invalid AI provider; please refer to the documentation for a list of valid providers"
)

// DescribeAction is the main action for the application.
func DescribeAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return ErrEmptyInput
	}

	var (
		key         = ctx.String("key")
		provider    = ctx.String("provider")
		model       = ctx.String("model")
		temperature = ctx.Float64("temperature")
		language    = ctx.String("language")
		context     = ctx.String("context")
		filename    = ctx.Bool("filename")
		image       = ctx.Args().Get(0)
		client      ai.Provider
	)

	if key == "" {
		return ErrEmptyKey
	}

	provider = strings.ToLower(provider)
	provider = strings.TrimSpace(provider)

	switch provider {
	case ai.ProviderOpenAI:
		client = openai.NewClient(key)

		if model == anthropic.DefaultModel {
			model = openai.DefaultModel
		}
	case ai.ProviderAnthropic:
		client = anthropic.NewClient(key)
	default:
		return ErrInvalidProvider
	}

	data, err := os.ReadFile(image)
	if err != nil {
		return fmt.Errorf("failed to read image: %w", err)
	}

	userPrompt := "Please generate a SEO-optimized alt text for the attached image."

	if language != "" {
		userPrompt += " User's preferred language: " + language
	}

	if context != "" {
		userPrompt += "\n\n Here is some context for the image: " + context
	}

	if filename {
		userPrompt += " Include a SEO-optimized filename as well."
	}

	req := ai.NewRequest(data, model, userPrompt, float32(temperature))

	err = client.Do(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}

	return nil
}
