package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Starts or continues a streaming Amazon Q Business conversation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_chatCmd).Standalone()

	qbusiness_chatCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application linked to a streaming Amazon Q Business conversation.")
	qbusiness_chatCmd.Flags().String("client-token", "", "A token that you provide to identify the chat input.")
	qbusiness_chatCmd.Flags().String("conversation-id", "", "The identifier of the Amazon Q Business conversation.")
	qbusiness_chatCmd.Flags().String("input-stream", "", "The streaming input for the `Chat` API.")
	qbusiness_chatCmd.Flags().String("parent-message-id", "", "The identifier used to associate a user message with a AI generated response.")
	qbusiness_chatCmd.Flags().String("user-groups", "", "The group names that a user associated with the chat input belongs to.")
	qbusiness_chatCmd.Flags().String("user-id", "", "The identifier of the user attached to the chat input.")
	qbusiness_chatCmd.MarkFlagRequired("application-id")
	qbusinessCmd.AddCommand(qbusiness_chatCmd)
}
