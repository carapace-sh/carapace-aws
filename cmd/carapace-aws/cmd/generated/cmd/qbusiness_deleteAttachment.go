package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteAttachmentCmd = &cobra.Command{
	Use:   "delete-attachment",
	Short: "Deletes an attachment associated with a specific Amazon Q Business conversation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_deleteAttachmentCmd).Standalone()

		qbusiness_deleteAttachmentCmd.Flags().String("application-id", "", "The unique identifier for the Amazon Q Business application environment.")
		qbusiness_deleteAttachmentCmd.Flags().String("attachment-id", "", "The unique identifier for the attachment.")
		qbusiness_deleteAttachmentCmd.Flags().String("conversation-id", "", "The unique identifier of the conversation.")
		qbusiness_deleteAttachmentCmd.Flags().String("user-id", "", "The unique identifier of the user involved in the conversation.")
		qbusiness_deleteAttachmentCmd.MarkFlagRequired("application-id")
		qbusiness_deleteAttachmentCmd.MarkFlagRequired("attachment-id")
		qbusiness_deleteAttachmentCmd.MarkFlagRequired("conversation-id")
	})
	qbusinessCmd.AddCommand(qbusiness_deleteAttachmentCmd)
}
