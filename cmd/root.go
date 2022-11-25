package cmd

import (
	"fmt"
	"os"

	"github.com/mskelton/gobble/config"
	"github.com/mskelton/gobble/rss"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobble",
	Short: "List RSS items from all feeds",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Read()

		for _, source := range cfg.Feeds {
			feed := rss.Read(source.Uri)

			fmt.Println("")
			fmt.Println(feed.Channel.Title)
			fmt.Println("--------------")

			for _, item := range feed.Channel.Items {
				fmt.Println(item.Title)
			}
		}
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
