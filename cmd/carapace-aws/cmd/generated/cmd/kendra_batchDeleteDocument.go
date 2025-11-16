package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_batchDeleteDocumentCmd = &cobra.Command{
	Use:   "batch-delete-document",
	Short: "Removes one or more documents from an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_batchDeleteDocumentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_batchDeleteDocumentCmd).Standalone()

		kendra_batchDeleteDocumentCmd.Flags().String("data-source-sync-job-metric-target", "", "")
		kendra_batchDeleteDocumentCmd.Flags().String("document-id-list", "", "One or more identifiers for documents to delete from the index.")
		kendra_batchDeleteDocumentCmd.Flags().String("index-id", "", "The identifier of the index that contains the documents to delete.")
		kendra_batchDeleteDocumentCmd.MarkFlagRequired("document-id-list")
		kendra_batchDeleteDocumentCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_batchDeleteDocumentCmd)
}
