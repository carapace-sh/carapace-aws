package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getAdmChannelCmd = &cobra.Command{
	Use:   "get-adm-channel",
	Short: "Retrieves information about the status and settings of the ADM channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getAdmChannelCmd).Standalone()

	pinpoint_getAdmChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getAdmChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_getAdmChannelCmd)
}
