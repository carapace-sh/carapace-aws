package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateGcmChannelCmd = &cobra.Command{
	Use:   "update-gcm-channel",
	Short: "Enables the GCM channel for an application or updates the status and settings of the GCM channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateGcmChannelCmd).Standalone()

	pinpoint_updateGcmChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_updateGcmChannelCmd.Flags().String("gcmchannel-request", "", "")
	pinpoint_updateGcmChannelCmd.MarkFlagRequired("application-id")
	pinpoint_updateGcmChannelCmd.MarkFlagRequired("gcmchannel-request")
	pinpointCmd.AddCommand(pinpoint_updateGcmChannelCmd)
}
