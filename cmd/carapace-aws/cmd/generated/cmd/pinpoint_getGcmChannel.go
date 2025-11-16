package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getGcmChannelCmd = &cobra.Command{
	Use:   "get-gcm-channel",
	Short: "Retrieves information about the status and settings of the GCM channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getGcmChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_getGcmChannelCmd).Standalone()

		pinpoint_getGcmChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_getGcmChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_getGcmChannelCmd)
}
