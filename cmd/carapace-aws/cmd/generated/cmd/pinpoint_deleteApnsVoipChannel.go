package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteApnsVoipChannelCmd = &cobra.Command{
	Use:   "delete-apns-voip-channel",
	Short: "Disables the APNs VoIP channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteApnsVoipChannelCmd).Standalone()

	pinpoint_deleteApnsVoipChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_deleteApnsVoipChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_deleteApnsVoipChannelCmd)
}
