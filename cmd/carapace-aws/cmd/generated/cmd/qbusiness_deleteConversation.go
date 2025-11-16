package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteConversationCmd = &cobra.Command{
	Use:   "delete-conversation",
	Short: "Deletes an Amazon Q Business web experience conversation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteConversationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_deleteConversationCmd).Standalone()

		qbusiness_deleteConversationCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application associated with the conversation.")
		qbusiness_deleteConversationCmd.Flags().String("conversation-id", "", "The identifier of the Amazon Q Business web experience conversation being deleted.")
		qbusiness_deleteConversationCmd.Flags().String("user-id", "", "The identifier of the user who is deleting the conversation.")
		qbusiness_deleteConversationCmd.MarkFlagRequired("application-id")
		qbusiness_deleteConversationCmd.MarkFlagRequired("conversation-id")
	})
	qbusinessCmd.AddCommand(qbusiness_deleteConversationCmd)
}
