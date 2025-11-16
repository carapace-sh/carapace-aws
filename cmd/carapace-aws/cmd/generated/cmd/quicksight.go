package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksightCmd = &cobra.Command{
	Use:   "quicksight",
	Short: "Amazon Quick Suite API Reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksightCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksightCmd).Standalone()

	})
	rootCmd.AddCommand(quicksightCmd)
}
