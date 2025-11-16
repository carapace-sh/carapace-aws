package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listAttachmentsCmd = &cobra.Command{
	Use:   "list-attachments",
	Short: "Gets a list of attachments associated with an Amazon Q Business web experience or a list of attachements associated with a specific Amazon Q Business conversation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listAttachmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listAttachmentsCmd).Standalone()

		qbusiness_listAttachmentsCmd.Flags().String("application-id", "", "The unique identifier for the Amazon Q Business application.")
		qbusiness_listAttachmentsCmd.Flags().String("conversation-id", "", "The unique identifier of the Amazon Q Business web experience conversation.")
		qbusiness_listAttachmentsCmd.Flags().String("max-results", "", "The maximum number of attachements to return.")
		qbusiness_listAttachmentsCmd.Flags().String("next-token", "", "If the number of attachments returned exceeds `maxResults`, Amazon Q Business returns a next token as a pagination token to retrieve the next set of attachments.")
		qbusiness_listAttachmentsCmd.Flags().String("user-id", "", "The unique identifier of the user involved in the Amazon Q Business web experience conversation.")
		qbusiness_listAttachmentsCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_listAttachmentsCmd)
}
