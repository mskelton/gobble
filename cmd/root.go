package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobble",
	Short: "List RSS items from all feeds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Bool("all", false, "Include items that have been read")
}
