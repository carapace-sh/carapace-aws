package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_sendUsersMessagesCmd = &cobra.Command{
	Use:   "send-users-messages",
	Short: "Creates and sends a message to a list of users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_sendUsersMessagesCmd).Standalone()

	pinpoint_sendUsersMessagesCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_sendUsersMessagesCmd.Flags().String("send-users-message-request", "", "")
	pinpoint_sendUsersMessagesCmd.MarkFlagRequired("application-id")
	pinpoint_sendUsersMessagesCmd.MarkFlagRequired("send-users-message-request")
	pinpointCmd.AddCommand(pinpoint_sendUsersMessagesCmd)
}
