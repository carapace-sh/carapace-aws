package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_sendMessageCmd = &cobra.Command{
	Use:   "send-message",
	Short: "Submits a message to the Amazon Q in Connect session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_sendMessageCmd).Standalone()

	qconnect_sendMessageCmd.Flags().String("assistant-id", "", "The identifier of the Amazon Q in Connect assistant.")
	qconnect_sendMessageCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_sendMessageCmd.Flags().String("configuration", "", "The configuration of the [SendMessage](https://docs.aws.amazon.com/connect/latest/APIReference/API_amazon-q-connect_SendMessage.html) request.")
	qconnect_sendMessageCmd.Flags().String("conversation-context", "", "The conversation context before the Amazon Q in Connect session.")
	qconnect_sendMessageCmd.Flags().String("message", "", "The message data to submit to the Amazon Q in Connect session.")
	qconnect_sendMessageCmd.Flags().String("session-id", "", "The identifier of the Amazon Q in Connect session.")
	qconnect_sendMessageCmd.Flags().String("type", "", "The message type.")
	qconnect_sendMessageCmd.MarkFlagRequired("assistant-id")
	qconnect_sendMessageCmd.MarkFlagRequired("message")
	qconnect_sendMessageCmd.MarkFlagRequired("session-id")
	qconnect_sendMessageCmd.MarkFlagRequired("type")
	qconnectCmd.AddCommand(qconnect_sendMessageCmd)
}
