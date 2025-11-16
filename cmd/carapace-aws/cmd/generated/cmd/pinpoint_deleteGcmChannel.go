package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteGcmChannelCmd = &cobra.Command{
	Use:   "delete-gcm-channel",
	Short: "Disables the GCM channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteGcmChannelCmd).Standalone()

	pinpoint_deleteGcmChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_deleteGcmChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_deleteGcmChannelCmd)
}
