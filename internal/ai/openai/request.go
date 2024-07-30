package openai

import (
	"git.sr.ht/~jamesponddotco/allalt/internal/ai"
	"git.sr.ht/~jamesponddotco/allalt/internal/prompt"
	"git.sr.ht/~jamesponddotco/allalt/internal/xbase64"
	"github.com/sashabaranov/go-openai"
)

// NewRequest creates a chat completion request with streaming support for the
// OpenAI API given the provided ai.Request object.
func NewRequest(req *ai.Request) openai.ChatCompletionRequest {
	image := xbase64.EncodeImageToDataURL(req.Image)

	return openai.ChatCompletionRequest{
		Model:       req.Model,
		Temperature: req.Temperature,
		Stream:      true,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: req.UserPrompt,
			},
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    image,
							Detail: openai.ImageURLDetailAuto,
						},
					},
				},
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt.System,
			},
		},
	}
}
