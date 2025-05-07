package ai

import (
	"context"

	"google.golang.org/genai"
)

type GeminiClient struct {
	client *genai.Client
}

func NewGeminiClient(apiKey string) *GeminiClient {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		panic(err)
	}

	return &GeminiClient{
		client: client,
	}
}

func (c *GeminiClient) GenerateCommitMessage(ctx context.Context, diff, language string) (string, error) {
	// Generate the prompt based on the language
	prompt := GeneratePrompt(language)
	model := "gemini-2.0-flash"

	// Create a new message request
	resp, err := c.client.Models.GenerateContent(ctx, model, genai.Text(prompt+diff), nil)

	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
