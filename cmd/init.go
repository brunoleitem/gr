package cmd

import (
	"github.com/brunoleitem/gr/handler"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// data.CreateTables()
		handler.Generate()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
