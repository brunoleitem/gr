package handler

import "github.com/brunoleitem/gr/internal/openai"

func Generate() {
	client := openai.CreateClient()
	client.Lala()
}
