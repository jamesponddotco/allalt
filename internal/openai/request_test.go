package openai_test

import (
	"strings"
	"testing"

	"git.sr.ht/~jamesponddotco/allalt/internal/openai"
)

func TestNewRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		language      string
		base64Image   string
		keywords      []string
		expectedModel string
	}{
		{
			name:          "English with keywords",
			language:      "English",
			base64Image:   "base64image1",
			keywords:      []string{"keyword1", "keyword2"},
			expectedModel: "gpt-4-vision-preview",
		},
		{
			name:          "Spanish without keywords",
			language:      "Spanish",
			base64Image:   "base64image2",
			keywords:      []string{},
			expectedModel: "gpt-4-vision-preview",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			request := openai.NewRequest(tt.language, tt.base64Image, tt.keywords)

			if request.MaxTokens <= 0 {
				t.Errorf("MaxTokens is non-positive, got %v", request.MaxTokens)
			}

			if request.Model != tt.expectedModel {
				t.Errorf("Model got %v, want %v", request.Model, tt.expectedModel)
			}

			if len(tt.keywords) > 0 && !strings.Contains(request.Messages[0].MultiContent[0].Text, strings.Join(tt.keywords, ", ")) {
				t.Errorf("Keywords not included properly in the system prompt")
			}
		})
	}
}
