package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getApnsVoipChannelCmd = &cobra.Command{
	Use:   "get-apns-voip-channel",
	Short: "Retrieves information about the status and settings of the APNs VoIP channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getApnsVoipChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getApnsVoipChannelCmd).Standalone()

		pinpoint_getApnsVoipChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getApnsVoipChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getApnsVoipChannelCmd)
}
