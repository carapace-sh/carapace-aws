package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateAdmChannelCmd = &cobra.Command{
	Use:   "update-adm-channel",
	Short: "Enables the ADM channel for an application or updates the status and settings of the ADM channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateAdmChannelCmd).Standalone()

	pinpoint_updateAdmChannelCmd.Flags().String("admchannel-request", "", "")
	pinpoint_updateAdmChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_updateAdmChannelCmd.MarkFlagRequired("admchannel-request")
	pinpoint_updateAdmChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_updateAdmChannelCmd)
}
