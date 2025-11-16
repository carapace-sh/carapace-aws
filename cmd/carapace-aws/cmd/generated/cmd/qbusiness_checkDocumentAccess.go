package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_checkDocumentAccessCmd = &cobra.Command{
	Use:   "check-document-access",
	Short: "Verifies if a user has access permissions for a specified document and returns the actual ACL attached to the document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_checkDocumentAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_checkDocumentAccessCmd).Standalone()

		qbusiness_checkDocumentAccessCmd.Flags().String("application-id", "", "The unique identifier of the application.")
		qbusiness_checkDocumentAccessCmd.Flags().String("data-source-id", "", "The unique identifier of the data source.")
		qbusiness_checkDocumentAccessCmd.Flags().String("document-id", "", "The unique identifier of the document.")
		qbusiness_checkDocumentAccessCmd.Flags().String("index-id", "", "The unique identifier of the index.")
		qbusiness_checkDocumentAccessCmd.Flags().String("user-id", "", "The unique identifier of the user.")
		qbusiness_checkDocumentAccessCmd.MarkFlagRequired("application-id")
		qbusiness_checkDocumentAccessCmd.MarkFlagRequired("document-id")
		qbusiness_checkDocumentAccessCmd.MarkFlagRequired("index-id")
		qbusiness_checkDocumentAccessCmd.MarkFlagRequired("user-id")
	})
	qbusinessCmd.AddCommand(qbusiness_checkDocumentAccessCmd)
}
