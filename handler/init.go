package handler

import (
	"github.com/brunoleitem/gr/internal/prompt"
	"github.com/brunoleitem/gr/models"
)

func Init() {
	pc := prompt.NewPrompt()
	askContent := pc.Create("Please provide a valid token", "OpenAPI Token: ")
	token := pc.Get(*askContent)
	cfg := models.Config{Key: token}
	cfg.Save()
}
