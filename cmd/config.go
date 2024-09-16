/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/brunoleitem/gr/data"
	"github.com/brunoleitem/gr/handler"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set up configuration file",
	Long:  `Set configuration files!`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTables()
		handler.Init()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
