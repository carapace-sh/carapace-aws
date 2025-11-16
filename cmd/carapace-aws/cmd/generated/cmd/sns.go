package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snsCmd = &cobra.Command{
	Use:   "sns",
	Short: "Amazon Simple Notification Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snsCmd).Standalone()

	})
	rootCmd.AddCommand(snsCmd)
}
