package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listMessagesCmd = &cobra.Command{
	Use:   "list-messages",
	Short: "Gets a list of messages associated with an Amazon Q Business web experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listMessagesCmd).Standalone()

	qbusiness_listMessagesCmd.Flags().String("application-id", "", "The identifier for the Amazon Q Business application.")
	qbusiness_listMessagesCmd.Flags().String("conversation-id", "", "The identifier of the Amazon Q Business web experience conversation.")
	qbusiness_listMessagesCmd.Flags().String("max-results", "", "The maximum number of messages to return.")
	qbusiness_listMessagesCmd.Flags().String("next-token", "", "If the number of messages returned exceeds `maxResults`, Amazon Q Business returns a next token as a pagination token to retrieve the next set of messages.")
	qbusiness_listMessagesCmd.Flags().String("user-id", "", "The identifier of the user involved in the Amazon Q Business web experience conversation.")
	qbusiness_listMessagesCmd.MarkFlagRequired("application-id")
	qbusiness_listMessagesCmd.MarkFlagRequired("conversation-id")
	qbusinessCmd.AddCommand(qbusiness_listMessagesCmd)
}
