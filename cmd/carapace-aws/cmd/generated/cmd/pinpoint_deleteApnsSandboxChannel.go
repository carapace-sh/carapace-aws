package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteApnsSandboxChannelCmd = &cobra.Command{
	Use:   "delete-apns-sandbox-channel",
	Short: "Disables the APNs sandbox channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteApnsSandboxChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteApnsSandboxChannelCmd).Standalone()

		pinpoint_deleteApnsSandboxChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteApnsSandboxChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteApnsSandboxChannelCmd)
}
