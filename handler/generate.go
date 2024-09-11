package handler

import (
	"github.com/brunoleitem/gr/internal/openai"
	"github.com/brunoleitem/gr/internal/prompt"
)

func Generate() {
	pc := prompt.NewPrompt()
	askContent := pc.Create("Please provide brief overview about your project!", "Overview: ")
	content := pc.Get(*askContent)
	client := openai.CreateClient()
	client.CreateReadme(content)
}
