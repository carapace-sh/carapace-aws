package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmailCmd = &cobra.Command{
	Use:   "pinpoint-email",
	Short: "Amazon Pinpoint Email Service\n\nWelcome to the *Amazon Pinpoint Email API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmailCmd).Standalone()

	})
	rootCmd.AddCommand(pinpointEmailCmd)
}
