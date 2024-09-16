package handler

import (
	"fmt"
	"os/exec"

	"github.com/brunoleitem/gr/internal/prompt"
)

func Generate() {
	pc := prompt.NewPrompt()
	askContent := pc.Create("Please provide brief overview about your project!", "Overview: ")
	content := pc.Get(*askContent)
	fmt.Println(content)

	cmd := exec.Command("git", "init")
	_, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	// client := openai.CreateClient()
	// client.CreateReadme(content)
}
