package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listDocumentMetadataHistoryCmd = &cobra.Command{
	Use:   "list-document-metadata-history",
	Short: "Amazon Web Services Systems Manager Change Manager will no longer be open to new customers starting November 7, 2025. If you would like to use Change Manager, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listDocumentMetadataHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listDocumentMetadataHistoryCmd).Standalone()

		ssm_listDocumentMetadataHistoryCmd.Flags().String("document-version", "", "The version of the change template.")
		ssm_listDocumentMetadataHistoryCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listDocumentMetadataHistoryCmd.Flags().String("metadata", "", "The type of data for which details are being requested.")
		ssm_listDocumentMetadataHistoryCmd.Flags().String("name", "", "The name of the change template.")
		ssm_listDocumentMetadataHistoryCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_listDocumentMetadataHistoryCmd.MarkFlagRequired("metadata")
		ssm_listDocumentMetadataHistoryCmd.MarkFlagRequired("name")
	})
	ssmCmd.AddCommand(ssm_listDocumentMetadataHistoryCmd)
}
