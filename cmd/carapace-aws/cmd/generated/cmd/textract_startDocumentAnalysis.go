package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_startDocumentAnalysisCmd = &cobra.Command{
	Use:   "start-document-analysis",
	Short: "Starts the asynchronous analysis of an input document for relationships between detected items such as key-value pairs, tables, and selection elements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_startDocumentAnalysisCmd).Standalone()

	textract_startDocumentAnalysisCmd.Flags().String("adapters-config", "", "Specifies the adapter to be used when analyzing a document.")
	textract_startDocumentAnalysisCmd.Flags().String("client-request-token", "", "The idempotent token that you use to identify the start request.")
	textract_startDocumentAnalysisCmd.Flags().String("document-location", "", "The location of the document to be processed.")
	textract_startDocumentAnalysisCmd.Flags().String("feature-types", "", "A list of the types of analysis to perform.")
	textract_startDocumentAnalysisCmd.Flags().String("job-tag", "", "An identifier that you specify that's included in the completion notification published to the Amazon SNS topic.")
	textract_startDocumentAnalysisCmd.Flags().String("kmskey-id", "", "The KMS key used to encrypt the inference results.")
	textract_startDocumentAnalysisCmd.Flags().String("notification-channel", "", "The Amazon SNS topic ARN that you want Amazon Textract to publish the completion status of the operation to.")
	textract_startDocumentAnalysisCmd.Flags().String("output-config", "", "Sets if the output will go to a customer defined bucket.")
	textract_startDocumentAnalysisCmd.Flags().String("queries-config", "", "")
	textract_startDocumentAnalysisCmd.MarkFlagRequired("document-location")
	textract_startDocumentAnalysisCmd.MarkFlagRequired("feature-types")
	textractCmd.AddCommand(textract_startDocumentAnalysisCmd)
}
