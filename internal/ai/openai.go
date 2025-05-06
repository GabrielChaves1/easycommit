package ai

import (
	"context"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	client := openai.NewClient(option.WithAPIKey(apiKey))
	return &OpenAIClient{
		client: &client,
	}
}

func (c *OpenAIClient) GenerateCommitMessage(ctx context.Context, diff, language string) (string, error) {
	prompt := GeneratePrompt(language)

	// Create a new chat completion request
	completion, err := c.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Model: openai.ChatModelGPT3_5Turbo,
			Messages: []openai.ChatCompletionMessageParamUnion{
				{
					OfSystem: &openai.ChatCompletionSystemMessageParam{
						Content: openai.ChatCompletionSystemMessageParamContentUnion{
							OfString: openai.String(prompt),
						},
					},
				},
				{
					OfUser: &openai.ChatCompletionUserMessageParam{
						Content: openai.ChatCompletionUserMessageParamContentUnion{
							OfString: openai.String(diff),
						},
					},
				},
			},
		})

	if err != nil {
		return "", err
	}

	// Check if the response contains choices
	if len(completion.Choices) == 0 {
		return "", nil
	}

	// Extract the content of the first choice
	content := completion.Choices[0].Message.Content

	// Check if the content is empty
	trimmedContent := strings.TrimSpace(content)

	// If the content is empty, return an empty string
	if trimmedContent == "" {
		return "", nil
	}

	return trimmedContent, nil
}

var _ Agent = &OpenAIClient{}
