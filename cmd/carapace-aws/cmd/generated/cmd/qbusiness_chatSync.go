package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_chatSyncCmd = &cobra.Command{
	Use:   "chat-sync",
	Short: "Starts or continues a non-streaming Amazon Q Business conversation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_chatSyncCmd).Standalone()

	qbusiness_chatSyncCmd.Flags().String("action-execution", "", "A request from an end user to perform an Amazon Q Business plugin action.")
	qbusiness_chatSyncCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application linked to the Amazon Q Business conversation.")
	qbusiness_chatSyncCmd.Flags().String("attachments", "", "A list of files uploaded directly during chat.")
	qbusiness_chatSyncCmd.Flags().String("attribute-filter", "", "Enables filtering of Amazon Q Business web experience responses based on document attributes or metadata fields.")
	qbusiness_chatSyncCmd.Flags().String("auth-challenge-response", "", "An authentication verification event response by a third party authentication server to Amazon Q Business.")
	qbusiness_chatSyncCmd.Flags().String("chat-mode", "", "The `chatMode` parameter determines the chat modes available to Amazon Q Business users:")
	qbusiness_chatSyncCmd.Flags().String("chat-mode-configuration", "", "The chat mode configuration for an Amazon Q Business application.")
	qbusiness_chatSyncCmd.Flags().String("client-token", "", "A token that you provide to identify a chat request.")
	qbusiness_chatSyncCmd.Flags().String("conversation-id", "", "The identifier of the Amazon Q Business conversation.")
	qbusiness_chatSyncCmd.Flags().String("parent-message-id", "", "The identifier of the previous system message in a conversation.")
	qbusiness_chatSyncCmd.Flags().String("user-groups", "", "The group names that a user associated with the chat input belongs to.")
	qbusiness_chatSyncCmd.Flags().String("user-id", "", "The identifier of the user attached to the chat input.")
	qbusiness_chatSyncCmd.Flags().String("user-message", "", "A end user message in a conversation.")
	qbusiness_chatSyncCmd.MarkFlagRequired("application-id")
	qbusinessCmd.AddCommand(qbusiness_chatSyncCmd)
}
