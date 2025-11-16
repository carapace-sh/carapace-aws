package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getApnsSandboxChannelCmd = &cobra.Command{
	Use:   "get-apns-sandbox-channel",
	Short: "Retrieves information about the status and settings of the APNs sandbox channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getApnsSandboxChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getApnsSandboxChannelCmd).Standalone()

		pinpoint_getApnsSandboxChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getApnsSandboxChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getApnsSandboxChannelCmd)
}
