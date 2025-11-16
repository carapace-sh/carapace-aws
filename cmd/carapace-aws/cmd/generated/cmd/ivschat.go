package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschatCmd = &cobra.Command{
	Use:   "ivschat",
	Short: "**Introduction**\n\nThe Amazon IVS Chat control-plane API enables you to create and manage Amazon IVS Chat resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschatCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschatCmd).Standalone()

	})
	rootCmd.AddCommand(ivschatCmd)
}
