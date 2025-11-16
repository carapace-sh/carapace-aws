package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getEmailChannelCmd = &cobra.Command{
	Use:   "get-email-channel",
	Short: "Retrieves information about the status and settings of the email channel for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getEmailChannelCmd).Standalone()

	pinpoint_getEmailChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getEmailChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_getEmailChannelCmd)
}
