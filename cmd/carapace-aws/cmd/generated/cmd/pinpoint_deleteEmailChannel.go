package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteEmailChannelCmd = &cobra.Command{
	Use:   "delete-email-channel",
	Short: "Disables the email channel for an application and deletes any existing settings for the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteEmailChannelCmd).Standalone()

	pinpoint_deleteEmailChannelCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_deleteEmailChannelCmd.MarkFlagRequired("application-id")
	pinpointCmd.AddCommand(pinpoint_deleteEmailChannelCmd)
}
