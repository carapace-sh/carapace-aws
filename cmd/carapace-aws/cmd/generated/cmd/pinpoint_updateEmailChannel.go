package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateEmailChannelCmd = &cobra.Command{
	Use:   "update-email-channel",
	Short: "Enables the email channel for an application or updates the status and settings of the email channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateEmailChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateEmailChannelCmd).Standalone()

		pinpoint_updateEmailChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateEmailChannelCmd.Flags().String("email-channel-request", "", "")
		pinpoint_updateEmailChannelCmd.MarkFlagRequired("application-id")
		pinpoint_updateEmailChannelCmd.MarkFlagRequired("email-channel-request")
	})
	pinpointCmd.AddCommand(pinpoint_updateEmailChannelCmd)
}
