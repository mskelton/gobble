package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mskelton/gobble/cache"
	"github.com/mskelton/gobble/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "gobble",
	Short:        "List RSS items from all feeds",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return err
		}

		if len(cfg.Feeds) == 0 {
			return errors.New("No feeds configured. Please configure a feed in ~/.config/gobble/config.yml")
		}

		c, err := cache.ReadS(cfg)
		if err != nil {
			return err
		}

		for i, feed := range c.Feeds {
			// Print a separator between feeds
			if i != 0 {
				fmt.Println()
			}

			// Print the feed title
			color.New(color.Bold, color.FgBlue).Println(feed.Title)

			for _, item := range feed.Items {
				fmt.Println(item.Title)
			}
		}

		return nil
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
