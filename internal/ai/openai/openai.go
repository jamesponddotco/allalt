// Package openai provides a client wrapper for the OpenAI API that complies
// with the ai.Provider interface.
package openai

import (
	"errors"
	"fmt"
	"io"

	"git.sr.ht/~jamesponddotco/allalt/internal/ai"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

// DefaultModel is the default model to use when making requests to the OpenAI
// API.
const DefaultModel = "gpt-4o"

// Client represents an OpenAI API client that complies with the ai.Provider
// interface.
type Client struct {
	// ai is the proper OpenAI API client.
	ai *openai.Client
}

// NewClient returns a new Client instance with the given API key.
func NewClient(key string) *Client {
	return &Client{
		ai: openai.NewClient(key),
	}
}

// Compile-time check to ensure Client implements the ai.Provider interface.
var _ ai.Provider = (*Client)(nil)

// Do performs a single API request to the OpenAI API, returning a response for
// the provided Request and writing said response to ctx.App.Writer as a stream
// of strings.
func (c *Client) Do(ctx *cli.Context, request *ai.Request) error {
	resp, err := c.ai.CreateChatCompletionStream(ctx.Context, NewRequest(request))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for {
		text, err := resp.Recv() //nolint:govet // Fixing this is more trouble than it's worth.
		if errors.Is(err, io.EOF) {
			fmt.Fprintf(ctx.App.Writer, "\n")

			break
		}

		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmt.Fprintf(ctx.App.Writer, "%s", text.Choices[0].Delta.Content)
	}

	return nil
}
