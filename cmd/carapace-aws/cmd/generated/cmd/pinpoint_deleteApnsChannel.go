package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteApnsChannelCmd = &cobra.Command{
	Use:   "delete-apns-channel",
	Short: "Disables the APNs channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteApnsChannelCmd).Standalone()

	pinpoint_deleteApnsChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_deleteApnsChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_deleteApnsChannelCmd)
}
