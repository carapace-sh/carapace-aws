package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateApnsChannelCmd = &cobra.Command{
	Use:   "update-apns-channel",
	Short: "Enables the APNs channel for an application or updates the status and settings of the APNs channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateApnsChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateApnsChannelCmd).Standalone()

		pinpoint_updateApnsChannelCmd.Flags().String("apnschannel-request", "", "")
		pinpoint_updateApnsChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateApnsChannelCmd.MarkFlagRequired("apnschannel-request")
		pinpoint_updateApnsChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_updateApnsChannelCmd)
}
