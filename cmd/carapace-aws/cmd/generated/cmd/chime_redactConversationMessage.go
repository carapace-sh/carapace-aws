package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_redactConversationMessageCmd = &cobra.Command{
	Use:   "redact-conversation-message",
	Short: "Redacts the specified message from the specified Amazon Chime conversation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_redactConversationMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_redactConversationMessageCmd).Standalone()

		chime_redactConversationMessageCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_redactConversationMessageCmd.Flags().String("conversation-id", "", "The conversation ID.")
		chime_redactConversationMessageCmd.Flags().String("message-id", "", "The message ID.")
		chime_redactConversationMessageCmd.MarkFlagRequired("account-id")
		chime_redactConversationMessageCmd.MarkFlagRequired("conversation-id")
		chime_redactConversationMessageCmd.MarkFlagRequired("message-id")
	})
	chimeCmd.AddCommand(chime_redactConversationMessageCmd)
}
