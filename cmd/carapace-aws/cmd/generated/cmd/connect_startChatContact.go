package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startChatContactCmd = &cobra.Command{
	Use:   "start-chat-contact",
	Short: "Initiates a flow to start a new chat for the customer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startChatContactCmd).Standalone()

	connect_startChatContactCmd.Flags().String("attributes", "", "A custom key-value pair using an attribute map.")
	connect_startChatContactCmd.Flags().String("chat-duration-in-minutes", "", "The total duration of the newly started chat session.")
	connect_startChatContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_startChatContactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow for initiating the chat.")
	connect_startChatContactCmd.Flags().String("customer-id", "", "The customer's identification number.")
	connect_startChatContactCmd.Flags().String("initial-message", "", "The initial message to be sent to the newly created chat.")
	connect_startChatContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_startChatContactCmd.Flags().String("participant-details", "", "Information identifying the participant.")
	connect_startChatContactCmd.Flags().String("persistent-chat", "", "Enable persistent chats.")
	connect_startChatContactCmd.Flags().String("related-contact-id", "", "The unique identifier for an Amazon Connect contact.")
	connect_startChatContactCmd.Flags().String("segment-attributes", "", "A set of system defined key-value pairs stored on individual contact segments using an attribute map.")
	connect_startChatContactCmd.Flags().String("supported-messaging-content-types", "", "The supported chat message content types.")
	connect_startChatContactCmd.MarkFlagRequired("contact-flow-id")
	connect_startChatContactCmd.MarkFlagRequired("instance-id")
	connect_startChatContactCmd.MarkFlagRequired("participant-details")
	connectCmd.AddCommand(connect_startChatContactCmd)
}
