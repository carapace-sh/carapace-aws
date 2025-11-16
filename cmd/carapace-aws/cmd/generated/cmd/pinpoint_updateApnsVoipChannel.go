package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateApnsVoipChannelCmd = &cobra.Command{
	Use:   "update-apns-voip-channel",
	Short: "Enables the APNs VoIP channel for an application or updates the status and settings of the APNs VoIP channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateApnsVoipChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateApnsVoipChannelCmd).Standalone()

		pinpoint_updateApnsVoipChannelCmd.Flags().String("apnsvoip-channel-request", "", "")
		pinpoint_updateApnsVoipChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateApnsVoipChannelCmd.MarkFlagRequired("apnsvoip-channel-request")
		pinpoint_updateApnsVoipChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_updateApnsVoipChannelCmd)
}
