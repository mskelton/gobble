package cmd

import (
	"github.com/mskelton/gobble/cache"
	"github.com/mskelton/gobble/config"
	"github.com/spf13/cobra"
)

var async bool

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get's the status of your unread feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.ReadConfig()
		if err != nil {
			return err
		}

		if async {
			// cache.Update()
		} else {
			cache.Sync(cfg)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().BoolVar(&async, "async", false, "Update the feeds in a background process")
}
