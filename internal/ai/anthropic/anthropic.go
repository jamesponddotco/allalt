// Package anthropic provides a client wrapper for the Anthropic API that
// complies with the ai.Provider interface.
package anthropic

import (
	"fmt"

	"git.sr.ht/~jamesponddotco/allalt/internal/ai"
	"github.com/liushuangls/go-anthropic/v2"
	"github.com/urfave/cli/v2"
)

// DefaultModel is the default model to use when making requests to the
// Anthropic API.
const DefaultModel = "claude-3-5-sonnet-20240620"

// Client represents an Anthropic API client that complies with the ai.Provider
// interface.
type Client struct {
	// ai is the proper Anthropic API client.
	ai *anthropic.Client
}

// Compile-time check to ensure Client implements the ai.Provider interface.
var _ ai.Provider = (*Client)(nil)

// NewClient returns a new Client instance with the given API key.
func NewClient(key string) *Client {
	return &Client{
		ai: anthropic.NewClient(key),
	}
}

// Do performs a single API request to the Anthropic API, returning a response
// for the provided Request and writing said response to ctx.App.Writer as a
// stream of strings.
func (c *Client) Do(ctx *cli.Context, request *ai.Request) error {
	_, err := c.ai.CreateMessagesStream(ctx.Context, NewRequest(ctx, request))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Fprintf(ctx.App.Writer, "\n")

	return nil
}
