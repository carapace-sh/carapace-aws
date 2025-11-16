package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsCmd = &cobra.Command{
	Use:   "ivs",
	Short: "**Introduction**",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsCmd).Standalone()

	})
	rootCmd.AddCommand(ivsCmd)
}
