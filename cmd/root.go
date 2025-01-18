package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "ggit",
	Short: "GGit is a simple Git implementation",
	Long:  "GGit is a simple Git implementation written in Go, designed to mimic Git functionality.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}
