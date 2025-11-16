package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_batchPutDocumentCmd = &cobra.Command{
	Use:   "batch-put-document",
	Short: "Adds one or more documents to an Amazon Q Business index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_batchPutDocumentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_batchPutDocumentCmd).Standalone()

		qbusiness_batchPutDocumentCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application.")
		qbusiness_batchPutDocumentCmd.Flags().String("data-source-sync-id", "", "The identifier of the data source sync during which the documents were added.")
		qbusiness_batchPutDocumentCmd.Flags().String("documents", "", "One or more documents to add to the index.")
		qbusiness_batchPutDocumentCmd.Flags().String("index-id", "", "The identifier of the Amazon Q Business index to add the documents to.")
		qbusiness_batchPutDocumentCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access your S3 bucket.")
		qbusiness_batchPutDocumentCmd.MarkFlagRequired("application-id")
		qbusiness_batchPutDocumentCmd.MarkFlagRequired("documents")
		qbusiness_batchPutDocumentCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_batchPutDocumentCmd)
}
