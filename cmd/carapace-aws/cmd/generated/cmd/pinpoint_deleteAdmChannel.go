package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteAdmChannelCmd = &cobra.Command{
	Use:   "delete-adm-channel",
	Short: "Disables the ADM channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteAdmChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_deleteAdmChannelCmd).Standalone()

		pinpoint_deleteAdmChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_deleteAdmChannelCmd.MarkFlagRequired("application-id")
	})
	pinpointCmd.AddCommand(pinpoint_deleteAdmChannelCmd)
}
