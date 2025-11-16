package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteSmsChannelCmd = &cobra.Command{
	Use:   "delete-sms-channel",
	Short: "Disables the SMS channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteSmsChannelCmd).Standalone()

	pinpoint_deleteSmsChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_deleteSmsChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_deleteSmsChannelCmd)
}
