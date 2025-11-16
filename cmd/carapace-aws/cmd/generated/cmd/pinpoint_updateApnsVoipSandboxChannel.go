package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateApnsVoipSandboxChannelCmd = &cobra.Command{
	Use:   "update-apns-voip-sandbox-channel",
	Short: "Enables the APNs VoIP sandbox channel for an application or updates the status and settings of the APNs VoIP sandbox channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateApnsVoipSandboxChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateApnsVoipSandboxChannelCmd).Standalone()

		pinpoint_updateApnsVoipSandboxChannelCmd.Flags().String("apnsvoip-sandbox-channel-request", "", "")
		pinpoint_updateApnsVoipSandboxChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateApnsVoipSandboxChannelCmd.MarkFlagRequired("apnsvoip-sandbox-channel-request")
		pinpoint_updateApnsVoipSandboxChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_updateApnsVoipSandboxChannelCmd)
}
