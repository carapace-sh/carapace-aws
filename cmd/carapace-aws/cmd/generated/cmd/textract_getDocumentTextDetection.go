package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_getDocumentTextDetectionCmd = &cobra.Command{
	Use:   "get-document-text-detection",
	Short: "Gets the results for an Amazon Textract asynchronous operation that detects text in a document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_getDocumentTextDetectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_getDocumentTextDetectionCmd).Standalone()

		textract_getDocumentTextDetectionCmd.Flags().String("job-id", "", "A unique identifier for the text detection job.")
		textract_getDocumentTextDetectionCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
		textract_getDocumentTextDetectionCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there are more blocks to retrieve), Amazon Textract returns a pagination token in the response.")
		textract_getDocumentTextDetectionCmd.MarkFlagRequired("job-id")
	})
	textractCmd.AddCommand(textract_getDocumentTextDetectionCmd)
}
