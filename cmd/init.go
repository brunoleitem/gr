package cmd

import (
	"github.com/brunoleitem/gr/data"
	"github.com/brunoleitem/gr/prompt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTables()
		setUserToken()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func setUserToken() {
	pc := prompt.NewPrompt()
	askContent := pc.Create("Please provide a valid token", "OpenAPI Token: ")
	token := pc.Get(*askContent)
	println(token)
}
