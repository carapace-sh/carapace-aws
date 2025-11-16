package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteApnsVoipSandboxChannelCmd = &cobra.Command{
	Use:   "delete-apns-voip-sandbox-channel",
	Short: "Disables the APNs VoIP sandbox channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteApnsVoipSandboxChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteApnsVoipSandboxChannelCmd).Standalone()

		pinpoint_deleteApnsVoipSandboxChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteApnsVoipSandboxChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteApnsVoipSandboxChannelCmd)
}
