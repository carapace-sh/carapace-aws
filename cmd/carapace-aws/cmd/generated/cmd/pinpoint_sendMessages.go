package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_sendMessagesCmd = &cobra.Command{
	Use:   "send-messages",
	Short: "Creates and sends a direct message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_sendMessagesCmd).Standalone()

	pinpoint_sendMessagesCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_sendMessagesCmd.Flags().String("message-request", "", "")
	pinpoint_sendMessagesCmd.MarkFlagRequired("application-id")
	pinpoint_sendMessagesCmd.MarkFlagRequired("message-request")
	pinpointCmd.AddCommand(pinpoint_sendMessagesCmd)
}
