package handler

import (
	"fmt"

	"github.com/brunoleitem/gr/internal/prompt"
)

func Generate() {
	pc := prompt.NewPrompt()
	askContent := pc.Create("Please provide brief overview about your project!", "Overview: ")
	content := pc.Get(*askContent)
	fmt.Println(content)
	// client := openai.CreateClient()
	// client.CreateReadme(content)
}
