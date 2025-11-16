package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getApnsChannelCmd = &cobra.Command{
	Use:   "get-apns-channel",
	Short: "Retrieves information about the status and settings of the APNs channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getApnsChannelCmd).Standalone()

	pinpoint_getApnsChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getApnsChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_getApnsChannelCmd)
}
