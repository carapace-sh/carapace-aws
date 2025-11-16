package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getApnsVoipSandboxChannelCmd = &cobra.Command{
	Use:   "get-apns-voip-sandbox-channel",
	Short: "Retrieves information about the status and settings of the APNs VoIP sandbox channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getApnsVoipSandboxChannelCmd).Standalone()

	pinpoint_getApnsVoipSandboxChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getApnsVoipSandboxChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_getApnsVoipSandboxChannelCmd)
}
