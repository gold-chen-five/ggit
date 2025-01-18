package cmd

import (
	"fmt"

	"github.com/gold-chen-five/ggit/internal/repo"
	"github.com/spf13/cobra"
)

// initCmd represents the 'init' command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new GGit repository",
	Long:  "The 'init' command initializes a new GGit repository by creating the necessary structure.",
	Run: func(cmd *cobra.Command, args []string) {
		err := repo.InitRepository()
		if err != nil {
			fmt.Printf("Error initializing repository: %v\n", err)
			return
		}
		fmt.Println("Initialized empty Git repository in .ggit/")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
