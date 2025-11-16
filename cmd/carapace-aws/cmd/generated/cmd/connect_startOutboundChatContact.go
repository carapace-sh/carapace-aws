package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startOutboundChatContactCmd = &cobra.Command{
	Use:   "start-outbound-chat-contact",
	Short: "Initiates a new outbound SMS contact to a customer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startOutboundChatContactCmd).Standalone()

	connect_startOutboundChatContactCmd.Flags().String("attributes", "", "A custom key-value pair using an attribute map.")
	connect_startOutboundChatContactCmd.Flags().String("chat-duration-in-minutes", "", "The total duration of the newly started chat session.")
	connect_startOutboundChatContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_startOutboundChatContactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow for the call.")
	connect_startOutboundChatContactCmd.Flags().String("destination-endpoint", "", "")
	connect_startOutboundChatContactCmd.Flags().String("initial-system-message", "", "")
	connect_startOutboundChatContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_startOutboundChatContactCmd.Flags().String("participant-details", "", "")
	connect_startOutboundChatContactCmd.Flags().String("related-contact-id", "", "The unique identifier for an Amazon Connect contact.")
	connect_startOutboundChatContactCmd.Flags().String("segment-attributes", "", "A set of system defined key-value pairs stored on individual contact segments using an attribute map.")
	connect_startOutboundChatContactCmd.Flags().String("source-endpoint", "", "")
	connect_startOutboundChatContactCmd.Flags().String("supported-messaging-content-types", "", "The supported chat message content types.")
	connect_startOutboundChatContactCmd.MarkFlagRequired("contact-flow-id")
	connect_startOutboundChatContactCmd.MarkFlagRequired("destination-endpoint")
	connect_startOutboundChatContactCmd.MarkFlagRequired("instance-id")
	connect_startOutboundChatContactCmd.MarkFlagRequired("segment-attributes")
	connect_startOutboundChatContactCmd.MarkFlagRequired("source-endpoint")
	connectCmd.AddCommand(connect_startOutboundChatContactCmd)
}
