/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gr",
	Short: "Git Init with AI",
	Long: `Init your git repository with a brand new markdown
  readme using AI.`,

	Run: func(cmd *cobra.Command, args []string) {
		// data.CreateTables()
		// handler.Generate()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
