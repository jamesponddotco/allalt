package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// Client represents an OpenAI API client.
type Client struct {
	// ai is the OpenAI API client.
	ai *openai.Client
}

// NewClient returns a new OpenAI API client.
func NewClient(key string) *Client {
	return &Client{
		ai: openai.NewClient(key),
	}
}

// Do performs the OpenAI API request.
func (c *Client) Do(ctx context.Context, req openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) { //nolint:gocritic // the underlying function doesn't take a pointer
	resp, err := c.ai.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}
