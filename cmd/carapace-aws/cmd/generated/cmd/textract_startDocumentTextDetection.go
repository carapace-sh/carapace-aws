package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_startDocumentTextDetectionCmd = &cobra.Command{
	Use:   "start-document-text-detection",
	Short: "Starts the asynchronous detection of text in a document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_startDocumentTextDetectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_startDocumentTextDetectionCmd).Standalone()

		textract_startDocumentTextDetectionCmd.Flags().String("client-request-token", "", "The idempotent token that's used to identify the start request.")
		textract_startDocumentTextDetectionCmd.Flags().String("document-location", "", "The location of the document to be processed.")
		textract_startDocumentTextDetectionCmd.Flags().String("job-tag", "", "An identifier that you specify that's included in the completion notification published to the Amazon SNS topic.")
		textract_startDocumentTextDetectionCmd.Flags().String("kmskey-id", "", "The KMS key used to encrypt the inference results.")
		textract_startDocumentTextDetectionCmd.Flags().String("notification-channel", "", "The Amazon SNS topic ARN that you want Amazon Textract to publish the completion status of the operation to.")
		textract_startDocumentTextDetectionCmd.Flags().String("output-config", "", "Sets if the output will go to a customer defined bucket.")
		textract_startDocumentTextDetectionCmd.MarkFlagRequired("document-location")
	})
	textractCmd.AddCommand(textract_startDocumentTextDetectionCmd)
}
