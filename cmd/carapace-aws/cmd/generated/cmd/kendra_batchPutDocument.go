package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_batchPutDocumentCmd = &cobra.Command{
	Use:   "batch-put-document",
	Short: "Adds one or more documents to an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_batchPutDocumentCmd).Standalone()

	kendra_batchPutDocumentCmd.Flags().String("custom-document-enrichment-configuration", "", "Configuration information for altering your document metadata and content during the document ingestion process when you use the `BatchPutDocument` API.")
	kendra_batchPutDocumentCmd.Flags().String("documents", "", "One or more documents to add to the index.")
	kendra_batchPutDocumentCmd.Flags().String("index-id", "", "The identifier of the index to add the documents to.")
	kendra_batchPutDocumentCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access your S3 bucket.")
	kendra_batchPutDocumentCmd.MarkFlagRequired("documents")
	kendra_batchPutDocumentCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_batchPutDocumentCmd)
}
