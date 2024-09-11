package openai

import (
	"context"

	"github.com/brunoleitem/gr/models"
	"github.com/sashabaranov/go-openai"
)

type OpenAiClient struct {
	client *openai.Client
}

func CreateClient() *OpenAiClient {
	cf := models.Config{}
	result, err := cf.Restore()
	if err != nil {
		panic(err)
	}

	client := openai.NewClient(result.Key)
	return &OpenAiClient{
		client: client,
	}
}

func (c *OpenAiClient) CreateReadme() {
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "Act like a specialist in development that create resumes to github readme with a quick description of the program",
				},
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "The resumes MUST be written in markdown and MUST be generic with technologies used and specific functions so the user can put manually.",
				},

			},
		}
	)
}
