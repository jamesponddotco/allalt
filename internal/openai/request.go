package openai

import (
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
	"github.com/sashabaranov/go-openai"
)

// _defaultMaxTokens is the default maximum number of tokens to generate.
const _defaultMaxTokens int = 100

// _systemPrompt is the system prompt to use when calling the OpenAI API.
const _systemPrompt string = `You are to adopt the expertise of Joost de Valk, a SEO and digital marketing expert. Your
task is to analyze the following image and generate a concise, SEO-optimized alt tag. Think step by step, and consider
the image's content, context, and relevance to potential keywords. Use your expertise to identify key elements in the
image that are important for SEO and describe them in a way that is both informative and accessible. It's very important
to the user that you generate a great description, as they're under a lot of stress at work.

Please generate an alt tag that improves the image's SEO performance. Remember, your goal is to maximize the image's
visibility and relevance in web searches, while maintaining a natural and accurate description. Don't output anything
else, just the value to the alt tag field. Do not use quotes and use a final period, just like the examples below.

Examples:
1. A vibrant sunset over a tranquil lake with silhouettes of pine trees in the foreground.
2. A bustling city street at night with illuminated skyscrapers.
3. Close-up of a colorful macaw perched on a tree branch in the rainforest.
4. Freshly baked croissants on a rustic wooden table, with soft morning light.

Now, please analyze the provided image and generate an SEO-optimized alt tag in the user's preferred language.
`

// NewRequest creates a chat completion request with streaming support for the OpenAI API.
func NewRequest(language, base64Image string, keywords []string) openai.ChatCompletionRequest {
	systemPrompt := _systemPrompt +
		"\n\n- User's preferred language: " + language

	if len(keywords) > 0 {
		systemPrompt += "\n- Keywords: " + xstrings.JoinWithSeparator(", ", keywords...)
	}

	return openai.ChatCompletionRequest{
		MaxTokens: _defaultMaxTokens,
		Model:     openai.GPT4VisionPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeText,
						Text: systemPrompt,
					},
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    base64Image,
							Detail: openai.ImageURLDetailAuto,
						},
					},
				},
			},
		},
		Stream: true,
	}
}
