// Package ai provides a generic interface for interacting with AI providers.
package ai

import "github.com/urfave/cli/v2"

// List of supported AI providers.
const (
	ProviderOpenAI    string = "openai"
	ProviderAnthropic string = "anthropic"
)

// Provider represents a provider like OpenAI, Anthropic, Ollama, etc.
type Provider interface {
	// Name returns the name of the provider.
	Name() string

	// Do performs a single API request to the provider's API, returning a
	// response for the provided Request.
	//
	// Do should attempt to interpret the response from the provider and return
	// it as streaming strings to ctx.App.Writer if successful.
	Do(ctx *cli.Context, req *Request) error
}

// Request represents an HTTP request to an AI provider's API.
type Request struct {
	// Context is a string that provides additional context for the request.
	Context string

	// Model is the full name of the LLM model to use for the request, e.g.
	// "gpt-4o-mini" or "claude-3-5-sonnet-20240620".
	Model string

	// UserPrompt is a second set of instructions for the AI model to follow
	// when generating a response.
	//
	// This prompt is sent to the AI provider after the system prompt and can be
	// used to provide additional context or constraints to the response.
	UserPrompt string

	// Image represents an image attachment that is sent to the AI provider
	// together with the user prompt.
	Image []byte

	// Temperature is a float between 0 and 1 that controls the randomness of
	// the response. A value of 0 will always return the most likely token,
	// while a value of 1 will sample from the distribution of tokens.
	Temperature float32
}

// NewRequest returns a new Request instance.
func NewRequest(image []byte, model, context, userPrompt string, temperature float32) *Request {
	return &Request{
		Image:       image,
		Model:       model,
		Context:     context,
		UserPrompt:  userPrompt,
		Temperature: temperature,
	}
}
