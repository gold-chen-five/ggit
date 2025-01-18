package cmd

import (
	"fmt"

	"github.com/gold-chen-five/ggit/internal/repo"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add the change file to staged area",
	Long:  "The 'add' command add the change file to staged area",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: No file specified")
			return
		}

		err := repo.AddFileToIndex(args[0])
		if err != nil {
			fmt.Println("Add file fial:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
