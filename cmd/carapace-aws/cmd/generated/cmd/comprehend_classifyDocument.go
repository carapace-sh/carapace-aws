package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_classifyDocumentCmd = &cobra.Command{
	Use:   "classify-document",
	Short: "Creates a classification request to analyze a single document in real-time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_classifyDocumentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_classifyDocumentCmd).Standalone()

		comprehend_classifyDocumentCmd.Flags().String("bytes", "", "Use the `Bytes` parameter to input a text, PDF, Word or image file.")
		comprehend_classifyDocumentCmd.Flags().String("document-reader-config", "", "Provides configuration parameters to override the default actions for extracting text from PDF documents and image files.")
		comprehend_classifyDocumentCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Number (ARN) of the endpoint.")
		comprehend_classifyDocumentCmd.Flags().String("text", "", "The document text to be analyzed.")
		comprehend_classifyDocumentCmd.MarkFlagRequired("endpoint-arn")
	})
	comprehendCmd.AddCommand(comprehend_classifyDocumentCmd)
}
