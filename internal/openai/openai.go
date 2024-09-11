package openai

import (
	"context"
	"fmt"

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
		panic("ERror to restore")
	}
	client := openai.NewClient(result.Key)
	return &OpenAiClient{
		client: client,
	}
}

func (c *OpenAiClient) CreateReadme(content string) {
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `
					You are an expert developer specializing in creating professional GitHub README summaries for projects.
					Your task is to generate a resume-like description of a GitHub project in Markdown format.
					The description MUST be concise, professional, and written in Markdown.
					`,
				},
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `
					The summary MUST include:
					- A table of the contents.
					- A brief overview of the project and its main features.
					- Technologies and tools used (leave placeholders for the user to fill in manually).
					- Major functions or features of the project (again, use placeholders for the user to specify).
					- A generic contributing guide
					- A default MIT license.
					DO NOT include specific implementation details or project-specific data; keep the summary generic and adaptable for any project.
					`,
				},
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `
					The structure should be formatted as a Markdown README template.
					`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
}
