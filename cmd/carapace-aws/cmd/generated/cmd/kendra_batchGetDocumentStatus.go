package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_batchGetDocumentStatusCmd = &cobra.Command{
	Use:   "batch-get-document-status",
	Short: "Returns the indexing status for one or more documents submitted with the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html) API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_batchGetDocumentStatusCmd).Standalone()

	kendra_batchGetDocumentStatusCmd.Flags().String("document-info-list", "", "A list of `DocumentInfo` objects that identify the documents for which to get the status.")
	kendra_batchGetDocumentStatusCmd.Flags().String("index-id", "", "The identifier of the index to add documents to.")
	kendra_batchGetDocumentStatusCmd.MarkFlagRequired("document-info-list")
	kendra_batchGetDocumentStatusCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_batchGetDocumentStatusCmd)
}
