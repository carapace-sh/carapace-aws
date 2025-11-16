package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snsCmd = &cobra.Command{
	Use:   "sns",
	Short: "Amazon Simple Notification Service\n\nAmazon Simple Notification Service (Amazon SNS) is a web service that enables you to build distributed web-enabled applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snsCmd).Standalone()

	})
	rootCmd.AddCommand(snsCmd)
}
