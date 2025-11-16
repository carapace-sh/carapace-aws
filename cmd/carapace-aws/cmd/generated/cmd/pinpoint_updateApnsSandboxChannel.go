package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateApnsSandboxChannelCmd = &cobra.Command{
	Use:   "update-apns-sandbox-channel",
	Short: "Enables the APNs sandbox channel for an application or updates the status and settings of the APNs sandbox channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateApnsSandboxChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_updateApnsSandboxChannelCmd).Standalone()

		pinpoint_updateApnsSandboxChannelCmd.Flags().String("apnssandbox-channel-request", "", "")
		pinpoint_updateApnsSandboxChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_updateApnsSandboxChannelCmd.MarkFlagRequired("apnssandbox-channel-request")
		pinpoint_updateApnsSandboxChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_updateApnsSandboxChannelCmd)
}
