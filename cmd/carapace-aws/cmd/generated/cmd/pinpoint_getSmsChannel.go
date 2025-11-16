package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getSmsChannelCmd = &cobra.Command{
	Use:   "get-sms-channel",
	Short: "Retrieves information about the status and settings of the SMS channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getSmsChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getSmsChannelCmd).Standalone()

		pinpoint_getSmsChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getSmsChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getSmsChannelCmd)
}
