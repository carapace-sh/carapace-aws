package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_batchDeleteDocumentCmd = &cobra.Command{
	Use:   "batch-delete-document",
	Short: "Asynchronously deletes one or more documents added using the `BatchPutDocument` API from an Amazon Q Business index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_batchDeleteDocumentCmd).Standalone()

	qbusiness_batchDeleteDocumentCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application.")
	qbusiness_batchDeleteDocumentCmd.Flags().String("data-source-sync-id", "", "The identifier of the data source sync during which the documents were deleted.")
	qbusiness_batchDeleteDocumentCmd.Flags().String("documents", "", "Documents deleted from the Amazon Q Business index.")
	qbusiness_batchDeleteDocumentCmd.Flags().String("index-id", "", "The identifier of the Amazon Q Business index that contains the documents to delete.")
	qbusiness_batchDeleteDocumentCmd.MarkFlagRequired("application-id")
	qbusiness_batchDeleteDocumentCmd.MarkFlagRequired("documents")
	qbusiness_batchDeleteDocumentCmd.MarkFlagRequired("index-id")
	qbusinessCmd.AddCommand(qbusiness_batchDeleteDocumentCmd)
}
