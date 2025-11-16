package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsCmd = &cobra.Command{
	Use:   "ivs",
	Short: "**Introduction**\n\nThe Amazon Interactive Video Service (IVS) API is REST compatible, using a standard HTTP API and an Amazon Web Services EventBridge event stream for responses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsCmd).Standalone()

	rootCmd.AddCommand(ivsCmd)
}
