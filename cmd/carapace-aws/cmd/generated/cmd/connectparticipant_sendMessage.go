package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_sendMessageCmd = &cobra.Command{
	Use:   "send-message",
	Short: "Sends a message.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_sendMessageCmd).Standalone()

	connectparticipant_sendMessageCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connectparticipant_sendMessageCmd.Flags().String("connection-token", "", "The authentication token associated with the connection.")
	connectparticipant_sendMessageCmd.Flags().String("content", "", "The content of the message.")
	connectparticipant_sendMessageCmd.Flags().String("content-type", "", "The type of the content.")
	connectparticipant_sendMessageCmd.MarkFlagRequired("connection-token")
	connectparticipant_sendMessageCmd.MarkFlagRequired("content")
	connectparticipant_sendMessageCmd.MarkFlagRequired("content-type")
	connectparticipantCmd.AddCommand(connectparticipant_sendMessageCmd)
}
