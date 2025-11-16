package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateSmsChannelCmd = &cobra.Command{
	Use:   "update-sms-channel",
	Short: "Enables the SMS channel for an application or updates the status and settings of the SMS channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateSmsChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateSmsChannelCmd).Standalone()

		pinpoint_updateSmsChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateSmsChannelCmd.Flags().String("smschannel-request", "", "")
		pinpoint_updateSmsChannelCmd.MarkFlagRequired("application-id")
		pinpoint_updateSmsChannelCmd.MarkFlagRequired("smschannel-request")
	})
	pinpointCmd.AddCommand(pinpoint_updateSmsChannelCmd)
}
