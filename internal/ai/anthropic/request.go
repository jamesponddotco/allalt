package anthropic

import (
	"fmt"

	"git.sr.ht/~jamesponddotco/allalt/internal/ai"
	"git.sr.ht/~jamesponddotco/allalt/internal/prompt"
	"github.com/liushuangls/go-anthropic/v2"
	"github.com/urfave/cli/v2"
)

// NewRequest creates a chat completion request with streaming support for the
// Anthropic API given the provided ai.Request object.
func NewRequest(ctx *cli.Context, req *ai.Request) anthropic.MessagesStreamRequest {
	return anthropic.MessagesStreamRequest{
		MessagesRequest: anthropic.MessagesRequest{
			Model:       req.Model,
			Temperature: &req.Temperature,
			System:      prompt.System,
			MaxTokens:   4096,
			Stream:      true,
			Messages: []anthropic.Message{
				{
					Role: anthropic.RoleUser,
					Content: []anthropic.MessageContent{
						anthropic.NewImageMessageContent(anthropic.MessageContentImageSource{
							Type:      "base64",
							MediaType: "image/jpeg",
							Data:      req.Image,
						}),
						anthropic.NewTextMessageContent(req.UserPrompt),
					},
				},
			},
		},
		OnContentBlockDelta: func(data anthropic.MessagesEventContentBlockDeltaData) {
			fmt.Fprintf(ctx.App.Writer, "%s", *data.Delta.Text)
		},
	}
}
