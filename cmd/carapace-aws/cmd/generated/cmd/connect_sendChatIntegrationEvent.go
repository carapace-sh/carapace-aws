package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_sendChatIntegrationEventCmd = &cobra.Command{
	Use:   "send-chat-integration-event",
	Short: "Processes chat integration events from Amazon Web Services or external integrations to Amazon Connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_sendChatIntegrationEventCmd).Standalone()

	connect_sendChatIntegrationEventCmd.Flags().String("destination-id", "", "Chat system identifier, used in part to uniquely identify chat.")
	connect_sendChatIntegrationEventCmd.Flags().String("event", "", "Chat integration event payload")
	connect_sendChatIntegrationEventCmd.Flags().String("new-session-details", "", "Contact properties to apply when starting a new chat.")
	connect_sendChatIntegrationEventCmd.Flags().String("source-id", "", "External identifier of chat customer participant, used in part to uniquely identify a chat.")
	connect_sendChatIntegrationEventCmd.Flags().String("subtype", "", "Classification of a channel.")
	connect_sendChatIntegrationEventCmd.MarkFlagRequired("destination-id")
	connect_sendChatIntegrationEventCmd.MarkFlagRequired("event")
	connect_sendChatIntegrationEventCmd.MarkFlagRequired("source-id")
	connectCmd.AddCommand(connect_sendChatIntegrationEventCmd)
}
